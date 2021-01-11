package adm

import (
	"github.com/paulmach/osm"
	"os"
	"testing"
)

func TestCreateIndex(t *testing.T) {
	os.MkdirAll("/run/shm/testdb", 0700)
	CreateIndex("/run/shm/ireland.osm.pbf", "/run/shm/testdb")
}

func TestEnumEntity(t *testing.T) {
	mapinde := GetMapFromDir("/run/shm/testdb")
	outo, oute := mapinde.ScanRegion(53.3532, -6.2598, 4)
	println("%v", len(outo))
	for _, v := range oute {
		println(v.String())
	}
}

func TestGetID(t *testing.T) {
	mapinde := GetMapFromDir("/run/shm/testdb")
	mapfile, err := os.Open("/run/shm/ireland.osm.pbf")
	if err != nil {
		panic(err)
	}
	outo := mapinde.GetFeatureByID("node/2718267438", mapfile)
	out := (*outo).(*osm.Node)
	out.FeatureID()
}

func TestGetRelation(t *testing.T) {
	mapinde := GetMapFromDir("/run/shm/testdb")
	outo := mapinde.GetRelationByFeature("node/990208736")
	for _, v := range outo {
		println(v.String())
	}

}

func TestGetRelation2(t *testing.T) {
	mapinde := GetMapFromDir("/run/shm/testdb")
	outo := mapinde.GetRelationByFeature("node/3451922931")
	for _, v := range outo {
		println(v.String())
	}

}

func TestSearch(t *testing.T) {
	mapinde := GetMapFromDir("/run/shm/testdb")
	outo, _ := mapinde.SearchByName("Sean MacDermott Street Lower")
	for _, v := range outo {
		println(v.String())
	}

}

func TestSearchPrefix(t *testing.T) {
	mapinde := GetMapFromDir("/run/shm/testdb")
	outo := mapinde.SearchByNamePrefix("Sean MacDermott")
	for _, v := range outo {
		println(v)
	}

}
