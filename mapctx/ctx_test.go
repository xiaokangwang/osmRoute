package mapctx

import (
	"github.com/beefsack/go-astar"
	"github.com/xiaokangwang/osmRoute/adm"
	"os"
	"testing"
)

type specDef struct {
}

func (s specDef) CanWalk() bool {
	return false
}

func (s specDef) CanDrive() bool {
	return true
}

func (s specDef) CanPublicTransport() bool {
	panic("implement me")
}

func TestMapCtx_ListRoutes(t *testing.T) {
	mapinde := adm.GetMapFromDir("/run/shm/testdb")
	mapfile, err := os.Open("/run/shm/ireland.osm.pbf")
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
	mapinde := adm.GetMapFromDir("/run/shm/testdb")
	mapfile, err := os.Open("/run/shm/ireland.osm.pbf")
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
	for _, v := range path {
		println(v.(*NodeImpl).FeatureID().String())
	}
}
