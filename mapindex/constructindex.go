package mapindex

import "github.com/paulmach/osm"

func ConstructIndexForObject(object osm.Object, skipLen int64) []MapIndex {
	ret := make([]MapIndex, 0, 6)

	{
		objid := MapID2ItemEntry{}
		objid.ID = object.ObjectID()
		objid.SkipLength = skipLen
		ret = append(ret, objid)
	}

	switch object.ObjectID().Type() {
	case osm.TypeNode:
		objNode := object.(*osm.Node)
		{
			objid := MapFeatureID2CurrentObjectIDEntry{}
			objid.ID = objNode.ObjectID()
			objid.IDFeat = objNode.FeatureID()
			ret = append(ret, objid)
		}
		{
			objloc := MapRegion2IDEntry{
				IDs:  []osm.ObjectID{objNode.ObjectID()},
				FIDs: []osm.FeatureID{objNode.FeatureID()},
				Lat:  objNode.Lat,
				Lon:  objNode.Lon,
			}
			ret = append(ret, objloc)
		}
	case osm.TypeWay:
		objWay := object.(*osm.Way)
		{
			objid := MapFeatureID2CurrentObjectIDEntry{}
			objid.ID = objWay.ObjectID()
			objid.IDFeat = objWay.FeatureID()
			ret = append(ret, objid)
		}
		{
			for _, v := range objWay.Nodes.FeatureIDs() {
				{
					objid := MapElementID2Refs{
						IDElem: v,
						Refs:   osm.FeatureIDs{objWay.FeatureID()},
					}
					ret = append(ret, objid)
				}
			}
		}
	case osm.TypeRelation:
		objRelation := object.(*osm.Relation)
		{
			objid := MapFeatureID2CurrentObjectIDEntry{}
			objid.ID = objRelation.ObjectID()
			objid.IDFeat = objRelation.FeatureID()
			ret = append(ret, objid)
		}
		{
			for _, v := range objRelation.Members.FeatureIDs() {
				{
					objid := MapElementID2Refs{
						IDElem: v,
						Refs:   osm.FeatureIDs{objRelation.FeatureID()},
					}
					ret = append(ret, objid)
				}
			}
		}
	}
	return ret
}
