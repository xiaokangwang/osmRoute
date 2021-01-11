package mapindex

import (
	"context"
	"github.com/dgraph-io/badger/v2"
	"github.com/paulmach/osm"
	"github.com/paulmach/osm/osmpbf"
	"io"
	"modernc.org/sortutil"
	"os"
	"strings"
)

func (m Map) GetFeatureByID(FeatureID string, file *os.File) *osm.Object {
	//Search First
	indexEntry := m.GetMapIndexByName("I_B_" + FeatureID)
	if indexEntry == nil {
		return nil
	}
	indexEV := indexEntry.(MapFeatureID2CurrentObjectIDEntry)
	_, err := file.Seek(indexEV.SkipLength, io.SeekStart)
	if err != nil {
		return nil
	}
	Scanner := osmpbf.New(context.Background(), file, 4)
	defer Scanner.Close()
	for Scanner.Scan() {
		object := Scanner.Object()
		if object.ObjectID() == indexEV.ID {
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
	Inte := tx.NewKeyIterator(Indexpfx, badger.IteratorOptions{
		PrefetchValues: true,
		AllVersions:    false,
		Prefix:         Indexpfx,
	})
	var ret []string
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

func (m Map) ScanRegion(Lat, Lon float64, mask int) (osm.ObjectIDs, osm.FeatureIDs) {
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
		Inte.Next()
	}
	Inte.Close()
	no := sortutil.Dedupe(ObjectIDSlice(Objs))
	Objs = Objs[:no]
	nf := sortutil.Dedupe(FeatureIDSlice(Feas))
	Feas = Feas[:nf]
	return Objs, Feas
}
