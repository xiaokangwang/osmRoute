package mapindex

import (
	"github.com/paulmach/osm"
	"github.com/xiaokangwang/osmRoute/interfacew"
)

func ConstructIndexForObjectSecondPass(object osm.Object, skipLen int64, m interfacew.MapResolver) []MapIndex {
	ret := make([]MapIndex, 0, 6)
	switch object.ObjectID().Type() {
	case osm.TypeRelation:
		objRelation := object.(*osm.Relation)
		routeval := objRelation.Tags.Find("route")
		if routeval == "bus" {
			for _, v := range objRelation.Members {
				switch v.Role {
				case "platform_entry_only":
					fallthrough
				case "platform":
					fallthrough
				case "platform_exit_only":
					{
						nodedata := (*m.ResolveInfoFromID(v.FeatureID().String())).(*osm.Node)

						objloc := MapRegion2IDEntry{
							Significant: []osm.FeatureID{objRelation.ID.FeatureID()},
							Lat:         nodedata.Lat,
							Lon:         nodedata.Lon,
						}
						ret = append(ret, objloc)
					}

				}
			}
		}
	}
	return ret
}
