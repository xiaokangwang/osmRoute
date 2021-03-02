package adm

import (
	"github.com/paulmach/osm"
	"github.com/stretchr/testify/assert"
	"github.com/xiaokangwang/osmRoute/admcommon"
	"github.com/xiaokangwang/osmRoute/util"
	"os"
	"testing"
)

func TestCreateIndex(t *testing.T) {
	os.MkdirAll(util.GetBaseDirFromEnvironment()+"/testdb", 0700)
	CreateIndex(util.GetBaseDirFromEnvironment()+"/ireland.osm.pbf", util.GetBaseDirFromEnvironment()+"/testdb")

	_, err := os.Stat(util.GetBaseDirFromEnvironment() + "/testdb")

	assert.Nil(t, err)
}

func TestEnumEntity(t *testing.T) {
	mapinde := admcommon.GetMapFromDir(util.GetBaseDirFromEnvironment() + "/testdb")
	outo, oute, _ := mapinde.ScanRegion(53.3532, -6.2598, 4)
	println("%v", len(outo))
	for _, v := range oute {
		println(v.String())
	}
}

func TestGetID(t *testing.T) {
	mapinde := admcommon.GetMapFromDir(util.GetBaseDirFromEnvironment() + "/testdb")
	mapfile, err := os.Open(util.GetBaseDirFromEnvironment() + "/ireland.osm.pbf")
	if err != nil {
		panic(err)
	}
	outo := mapinde.GetFeatureByID("node/2718267438", mapfile)
	out := (*outo).(*osm.Node)
	out.FeatureID()
	assert.Equal(t, "node/2718267438", out.FeatureID().String())
}

func TestGetID2(t *testing.T) {
	mapinde := admcommon.GetMapFromDir(util.GetBaseDirFromEnvironment() + "/testdb")
	mapfile, err := os.Open(util.GetBaseDirFromEnvironment() + "/ireland.osm.pbf")
	if err != nil {
		panic(err)
	}
	outo := mapinde.GetFeatureByID("way/52879958", mapfile)
	out := (*outo).(*osm.Way)
	out.FeatureID()
	assert.Equal(t, "way/52879958", out.FeatureID().String())
}

func TestGetRelation(t *testing.T) {
	mapinde := admcommon.GetMapFromDir(util.GetBaseDirFromEnvironment() + "/testdb")
	outo := mapinde.GetRelationByFeature("node/990208736")
	for _, v := range outo {
		println(v.String())
	}

}

func TestGetRelation2(t *testing.T) {
	mapinde := admcommon.GetMapFromDir(util.GetBaseDirFromEnvironment() + "/testdb")
	outo := mapinde.GetRelationByFeature("node/3451922931")
	for _, v := range outo {
		println(v.String())
	}

}

func TestSearch(t *testing.T) {
	mapinde := admcommon.GetMapFromDir(util.GetBaseDirFromEnvironment() + "/testdb")
	outo, _ := mapinde.SearchByName("Sean MacDermott Street Lower")
	for _, v := range outo {
		println(v.String())
	}
	assert.GreaterOrEqual(t, 1, len(outo))
}

func TestSearchPrefix(t *testing.T) {
	mapinde := admcommon.GetMapFromDir(util.GetBaseDirFromEnvironment() + "/testdb")
	outo := mapinde.SearchByNamePrefix("Sean MacDermott")
	for _, v := range outo {
		println(v)
	}
	assert.GreaterOrEqual(t, 1, len(outo))
}
