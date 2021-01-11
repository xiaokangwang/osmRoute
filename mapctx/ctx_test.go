package mapctx

import (
	"github.com/xiaokangwang/osmRoute/adm"
	"os"
	"testing"
)

type specDef struct {
}

func (s specDef) CanWalk() bool {
	return true
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
