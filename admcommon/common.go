package admcommon

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/sirupsen/logrus"
	"github.com/xiaokangwang/osmRoute/mapindex"
)

func GetMapFromDir(dbDir string) *mapindex.Map {
	logger := logrus.WithField("module", "database")
	opts := badger.DefaultOptions(dbDir)
	opts.Logger = logger
	db, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}
	newMap := mapindex.NewMap(db)
	return newMap
}
