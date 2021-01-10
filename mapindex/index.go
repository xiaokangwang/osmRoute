package mapindex

import (
	"context"
	"github.com/dgraph-io/badger/v2"
	"github.com/paulmach/osm/osmpbf"
	"os"
)

type Map struct {
	db *badger.DB
}

func (m *Map) ConstructIndex(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	Scanner := osmpbf.New(context.Background(), file, 4)
	for {
		object := Scanner.Object()
		fsb := Scanner.FullyScannedBytes()
		indexs := ConstructIndexForObject(object, fsb)
		tx := m.db.NewTransaction(true)
		for _, v := range indexs {
			var mapIndex MapIndex
			key := v.ToIndexKey()
			keyval, err := tx.Get([]byte(key))
			if err == nil {
				_ = keyval.Value(func(val []byte) error {
					mapIndex = Load(key, val)
					return nil
				})
			}
			newval := v.Install(mapIndex)
			_, newvalBinary := Store(newval)
			err = tx.Set([]byte(key), newvalBinary)
			if err != nil {
				panic(err)
			}
		}
		err = tx.Commit()
		if err != nil {
			panic(err)
		}
		tx.Discard()
		if !Scanner.Scan() {
			break
		}
	}
	return nil
}

func NewMap(db *badger.DB) *Map {
	return &Map{db: db}
}
