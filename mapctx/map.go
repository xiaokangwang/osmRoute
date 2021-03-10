package mapctx

import (
	"github.com/beefsack/go-astar"
	"github.com/paulmach/osm"
)

type Node interface {
	astar.Pather
	FindConnection(ConnectionSpec) []Connection
	PathNeighborVia(to astar.Pather) osm.Object
}

type Connection interface {
	From() Node
	To() Node
	GetCost() float64
	Via() osm.Object
}

type ConnectionSpec interface {
	CanWalk() bool
	CanDrive() bool
	CanPublicTransport() bool

	TimeFactor() float64
	CostFactor() float64
	SustainableFactor() float64
}
