package mapindex

import (
	"github.com/dgraph-io/badger/v2"
	"github.com/paulmach/osm"
	"modernc.org/sortutil"
)

func (m Map) GetObjectByID(ObjectID string) {

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
