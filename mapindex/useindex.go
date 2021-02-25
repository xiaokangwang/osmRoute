package mapindex

import (
	"context"
	"github.com/dgraph-io/badger/v3"
	"github.com/fxamacker/cbor"
	"github.com/paulmach/osm"
	"github.com/paulmach/osm/osmpbf"
	"io"
	"modernc.org/sortutil"
	"os"
	"sort"
	"strings"
	"time"
)

func (m Map) GetFeatureByID(FeatureID string, file *os.File) *osm.Object {

	//See if we can hit the cache
	obj, found := m.objCache.Get(FeatureID)
	if found {
		return obj.(*osm.Object)
	}

	//Search First
	indexEntry := m.GetMapIndexByName("I_B_" + FeatureID)
	if indexEntry == nil {
		return nil
	}
	indexEV := indexEntry.(MapFeatureID2CurrentObjectIDEntry)

	if indexEV.CachedData != nil && len(indexEV.CachedData) != 0 {
		var objw osm.Object
		switch indexEV.IDFeat.Type() {
		case osm.TypeNode:
			var node osm.Node
			err := cbor.Unmarshal(indexEV.CachedData, &node)
			if err != nil {
				panic(err)
			}
			s := &node
			objw = s
		case osm.TypeWay:
			var way osm.Way
			err := cbor.Unmarshal(indexEV.CachedData, &way)
			if err != nil {
				panic(err)
			}
			s := &way
			objw = s

		case osm.TypeRelation:
			var rela osm.Relation
			err := cbor.Unmarshal(indexEV.CachedData, &rela)
			if err != nil {
				panic(err)
			}
			s := &rela
			objw = s
		}
		return &objw
	}

	_, err := file.Seek(indexEV.SkipLength, io.SeekStart)
	if err != nil {
		return nil
	}
	Scanner := osmpbf.New(context.Background(), file, 4)
	defer Scanner.Close()
	for Scanner.Scan() {
		object := Scanner.Object()

		var feaid osm.FeatureID
		switch object.ObjectID().Type() {
		case osm.TypeNode:
			objNode := object.(*osm.Node)

			feaid = objNode.FeatureID()

		case osm.TypeWay:
			objWay := object.(*osm.Way)

			feaid = objWay.FeatureID()
		case osm.TypeRelation:
			objRelation := object.(*osm.Relation)
			feaid = objRelation.FeatureID()
		default:
			panic("Unexpected Type")
		}

		m.objCache.Set(feaid.String(), &object, time.Second*30)

		if object.ObjectID() == indexEV.ID {
			var cacheByte []byte

			switch object.ObjectID().Type() {
			case osm.TypeNode:
				objNode := object.(*osm.Node)

				byte, err := cbor.Marshal(*objNode, cbor.CanonicalEncOptions())
				if err != nil {
					panic(err)
				}
				cacheByte = byte
			case osm.TypeWay:
				objWay := object.(*osm.Way)

				byte, err := cbor.Marshal(*objWay, cbor.CanonicalEncOptions())
				if err != nil {
					panic(err)
				}
				cacheByte = byte
			case osm.TypeRelation:
				objRelation := object.(*osm.Relation)

				byte, err := cbor.Marshal(*objRelation, cbor.CanonicalEncOptions())
				if err != nil {
					panic(err)
				}
				cacheByte = byte
			default:
				panic("Unexpected Type")
			}

			tx := m.db.NewTransaction(true)

			indexEV.CachedData = cacheByte

			indexEVWriteback, err := cbor.Marshal(indexEV, cbor.CanonicalEncOptions())
			if err != nil {
				panic(err)
			}
			err = tx.Set([]byte(indexEV.ToIndexKey()), indexEVWriteback)
			if err != nil {
				panic(err)
			}
			err = tx.Commit()
			if err != nil {
				panic(err)
			}
			tx.Discard()
			retObj := Scanner.Object()
			return &retObj
		}
	}
	return nil
}

func (m Map) GetRelationByFeature(FeatureID string) osm.FeatureIDs {
	indexEntry := m.GetMapIndexByName("I_D_" + FeatureID)
	if indexEntry == nil {
		return nil
	}
	indexEV := indexEntry.(MapFeatureID2Refs)
	sort.Sort(sort.Reverse(FeatureIDSlice(indexEV.Refs)))
	return indexEV.Refs
}

func (m Map) SearchByName(Name string) (osm.FeatureIDs, uint32) {
	indexEntry := m.GetMapIndexByName("I_E_" + Name)
	if indexEntry == nil {
		return nil, 0
	}
	indexEV := indexEntry.(MapFTS)
	return indexEV.Refs, indexEV.Count
}

func (m Map) SearchByNamePrefix(Name string) []string {
	perfix := "I_E_" + Name
	Indexpfx := []byte(perfix)
	tx := m.db.NewTransaction(false)
	defer tx.Discard()
	Inte := tx.NewIterator(badger.IteratorOptions{
		PrefetchValues: true,
		AllVersions:    false,
		Prefix:         Indexpfx,
	})
	var ret []string
	Inte.Seek(Indexpfx)
	for {
		if !Inte.ValidForPrefix(Indexpfx) {
			break
		}
		key := string(Inte.Item().Key())
		outs := strings.SplitN(key, "_", 3)
		ret = append(ret, outs[2])
		if len(ret) >= 30 {
			break
		}
		Inte.Next()
	}
	Inte.Close()
	return ret
}

func (m Map) GetMapIndexByName(name string) MapIndex {
	tx := m.db.NewTransaction(false)
	defer tx.Discard()
	item, err := tx.Get([]byte(name))
	if err != nil {
		return nil
	}
	var mapIndex MapIndex
	_ = item.Value(func(val []byte) error {
		mapIndex = Load(name, val)
		return nil
	})
	return mapIndex
}

func (m Map) ScanRegion(Lat, Lon float64, mask int) (osm.ObjectIDs, osm.FeatureIDs, []FeatureIDWithLocation) {
	return m.scanRegion(Lat, Lon, mask)
}

func (m Map) scanRegion(Lat float64, Lon float64, mask int) (osm.ObjectIDs, osm.FeatureIDs, []FeatureIDWithLocation) {
	Index := "I_C_" + InterlacedString(FloatToConstantLengthFloat(Lat), FloatToConstantLengthFloat(Lon))
	Index = Index[:len(Index)-mask*2]
	Indexpfx := []byte(Index)
	tx := m.db.NewTransaction(false)
	defer tx.Discard()
	Inte := tx.NewIterator(badger.IteratorOptions{
		PrefetchValues: true,
		AllVersions:    false,
		Prefix:         Indexpfx,
	})
	Objs := make(osm.ObjectIDs, 0, 16)
	Feas := make(osm.FeatureIDs, 0, 16)
	FeasP := make([]FeatureIDWithLocation, 0, 16)
	Inte.Rewind()
	for {
		if !Inte.ValidForPrefix(Indexpfx) {
			break
		}
		var mapindex MapIndex
		_ = Inte.Item().Value(func(val []byte) error {
			mapindex = Load(string(Inte.Item().Key()), val)
			return nil
		})
		index := mapindex.(MapRegion2IDEntry)
		Objs = append(Objs, index.IDs...)
		Feas = append(Feas, index.FIDs...)

		fesl := FeatureIDWithLocation{
			Lat:                  index.Lat,
			Lon:                  index.Lon,
			FeatureIDs:           index.FIDs,
			SignificantRelations: index.Significant,
		}
		FeasP = append(FeasP, fesl)
		Inte.Next()
	}
	Inte.Close()
	no := sortutil.Dedupe(ObjectIDSlice(Objs))
	Objs = Objs[:no]
	nf := sortutil.Dedupe(FeatureIDSlice(Feas))
	Feas = Feas[:nf]
	return Objs, Feas, FeasP
}

type FeatureIDWithLocation struct {
	osm.FeatureIDs
	Lat float64
	Lon float64

	SignificantRelations osm.FeatureIDs
}
