package mapindex

import (
	"context"
	"github.com/dgraph-io/badger/v3"
	"github.com/patrickmn/go-cache"
	"github.com/paulmach/osm/osmpbf"
	"github.com/xiaokangwang/osmRoute/interfacew"
	"os"
	"time"
)

type Map struct {
	db       *badger.DB
	objCache *cache.Cache
}

func (m *Map) ConstructIndex(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	Scanner := osmpbf.New(context.Background(), file, 4)
	defer Scanner.Close()
	for Scanner.Scan() {
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
	}
	return nil
}

func (m *Map) ConstructIndexSecondPass(path string, resolver interfacew.MapResolver) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	Scanner := osmpbf.New(context.Background(), file, 4)
	defer Scanner.Close()
	for Scanner.Scan() {
		object := Scanner.Object()
		fsb := Scanner.FullyScannedBytes()
		indexs := ConstructIndexForObjectSecondPass(object, fsb, resolver)
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
	}
	return nil
}

func NewMap(db *badger.DB) *Map {
	return &Map{db: db, objCache: cache.New(time.Second*30, time.Second*60)}
}
