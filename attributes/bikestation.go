package attributes

import "github.com/xiaokangwang/osmRoute/interfacew"

type BikeStation struct {
}

func (b *BikeStation) ListAllRoutes() []interfacew.MapLocation {
	return []interfacew.MapLocation{}
}
