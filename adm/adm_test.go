package adm

import (
	"os"
	"testing"
)

func TestCreateIndex(t *testing.T) {
	os.MkdirAll("/run/shm/testdb", 0700)
	CreateIndex("/run/shm/ireland.osm.pbf", "/run/shm/testdb")
}

func TestEnumEntity(t *testing.T) {
	mapinde := GetMapFromDir("/run/shm/testdb")
	outo, oute := mapinde.ScanRegion(53.35342, -6.26233, 3)
	println("%v", len(outo))
	for _, v := range oute {
		println(v.String())
	}
}
