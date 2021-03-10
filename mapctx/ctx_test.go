package mapctx

import (
	"github.com/beefsack/go-astar"
	"github.com/paulmach/osm"
	"github.com/stretchr/testify/assert"
	"github.com/xiaokangwang/osmRoute/admcommon"
	"github.com/xiaokangwang/osmRoute/util"
	"os"
	"reflect"
	"testing"
)

type specDef struct {
}

func (s specDef) TimeFactor() float64 {
	return 1
}

func (s specDef) CostFactor() float64 {
	return 1
}

func (s specDef) SustainableFactor() float64 {
	return 1
}

func (s specDef) CanWalk() bool {
	return false
}

func (s specDef) CanDrive() bool {
	return true
}

func (s specDef) CanPublicTransport() bool {
	return true
}

func TestMapCtx_ListRoutes(t *testing.T) {
	mapinde := admcommon.GetMapFromDir(util.GetBaseDirFromEnvironment() + "/testdb")
	mapfile, err := os.Open(util.GetBaseDirFromEnvironment() + "/ireland.osm.pbf")
	if err != nil {
		panic(err)
	}
	mapCtx := NewMapCtx(*mapinde, mapfile)

	InitialNodes := mapCtx.GetNodeWithInterconnection(53.3532, -6.2598, specDef{})

	for _, v := range InitialNodes {
		node := v.(*NodeImpl)
		println(node.FeatureID().String())
	}
}

func TestMapCtx_TryRoute(t *testing.T) {
	mapinde := admcommon.GetMapFromDir(util.GetBaseDirFromEnvironment() + "/testdb")
	mapfile, err := os.Open(util.GetBaseDirFromEnvironment() + "/ireland.osm.pbf")
	if err != nil {
		panic(err)
	}
	mapCtx := NewMapCtx(*mapinde, mapfile)

	mapCtx.SetSpec(specDef{})

	InitialNodes := mapCtx.GetNodeWithInterconnection(53.35214, -6.25866, specDef{})

	for _, v := range InitialNodes {
		node := v.(*NodeImpl)
		println(node.FeatureID().String())
	}

	InitialNodesF := mapCtx.GetNodeWithInterconnection(53.36135, -6.23813, specDef{})

	for _, v := range InitialNodesF {
		node := v.(*NodeImpl)
		println(node.FeatureID().String())
	}
	path, dist, found := astar.Path(InitialNodes[0], InitialNodesF[0])
	println(found)
	println(dist)
	var last astar.Pather
	reverseAny(path)
	for _, v := range path {
		if last != nil {
			fid := ""
			ViaObject := last.(Node).PathNeighborVia(v)
			viatype := ViaObject.ObjectID().Type()
			switch viatype {
			case osm.TypeWay:
				infoway := (ViaObject).(*osm.Way)
				fid = infoway.FeatureID().String()
			case osm.TypeRelation:
				inforela := (ViaObject).(*osm.Relation)
				_ = inforela
				fid = inforela.FeatureID().String()
			}
			println("via:", fid)
		}
		println(v.(*NodeImpl).FeatureID().String())
		last = v
	}
}

func TestMapCtx_TryRouteWithBus(t *testing.T) {
	mapinde := admcommon.GetMapFromDir(util.GetBaseDirFromEnvironment() + "/testdb")
	mapfile, err := os.Open(util.GetBaseDirFromEnvironment() + "/ireland.osm.pbf")
	if err != nil {
		panic(err)
	}
	mapCtx := NewMapCtx(*mapinde, mapfile)

	mapCtx.SetSpec(specDef{})

	InitialNodes := mapCtx.GetNodeWithInterconnection(53.3630144, -6.2584680, specDef{})

	for _, v := range InitialNodes {
		node := v.(*NodeImpl)
		println(node.FeatureID().String())
	}

	InitialNodesF := mapCtx.GetNodeWithInterconnection(53.3918925, -6.2463432, specDef{})

	for _, v := range InitialNodesF {
		node := v.(*NodeImpl)
		println(node.FeatureID().String())
	}
	path, dist, found := astar.Path(InitialNodes[0], InitialNodesF[0])
	println(found)
	println(dist)
	var last astar.Pather
	reverseAny(path)
	ok := false
	for _, v := range path {
		if last != nil {
			fid := ""
			ViaObject := last.(Node).PathNeighborVia(v)
			viatype := ViaObject.ObjectID().Type()
			switch viatype {
			case osm.TypeWay:
				infoway := (ViaObject).(*osm.Way)
				fid = infoway.FeatureID().String()
			case osm.TypeRelation:
				inforela := (ViaObject).(*osm.Relation)
				if inforela.Tags.Find("route") == "bus" {
					ok = true
				}
				_ = inforela
				fid = inforela.FeatureID().String()
			}
			println("via:", fid)
		}
		println(v.(*NodeImpl).FeatureID().String())
		last = v
	}
	assert.True(t, ok, "Should have at least a bus route")
}

func reverseAny(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
