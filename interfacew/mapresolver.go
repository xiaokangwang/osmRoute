package interfacew

import "github.com/paulmach/osm"

type MapResolver interface {
	ResolveInfoFromID(FeaID string) *osm.Object
}

type MapLocation struct {
	Lat, Lon float64
}
