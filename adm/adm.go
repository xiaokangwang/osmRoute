package adm

import (
	"github.com/dgraph-io/badger/v3"
	log "github.com/sirupsen/logrus"
	"github.com/xiaokangwang/osmRoute/mapctx"
	"github.com/xiaokangwang/osmRoute/mapindex"
	"github.com/xiaokangwang/osmRoute/util"
	"os"
)

func CreateIndex(mapPath string, dbDir string) {
	logger := log.WithField("module", "database")
	opts := badger.DefaultOptions(dbDir)
	opts.Logger = logger
	db, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}
	CreateIndexDb(mapPath, db)
}

func CreateIndexDb(mapPath string, db *badger.DB) {
	newMap := mapindex.NewMap(db)
	_ = newMap.ConstructIndex(mapPath)

	mapfile, err := os.Open(util.GetBaseDirFromEnvironment() + "/ireland.osm.pbf")
	if err != nil {
		panic(err)
	}
	mapCtx := mapctx.NewMapCtx(*newMap, mapfile)

	_ = newMap.ConstructIndexSecondPass(mapPath, mapCtx)
}
