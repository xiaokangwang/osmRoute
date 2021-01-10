package adm

import (
	"github.com/dgraph-io/badger/v2"
	"github.com/xiaokangwang/osmRoute/mapindex"
)

func CreateIndex(mapPath string, dbDir string) {
	db, err := badger.Open(badger.DefaultOptions(dbDir))
	if err != nil {
		panic(err)
	}
	CreateIndexDb(mapPath, db)
}

func CreateIndexDb(mapPath string, db *badger.DB) {
	newMap := mapindex.NewMap(db)
	_ = newMap.ConstructIndex(mapPath)
}
