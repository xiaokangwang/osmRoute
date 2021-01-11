package mapctx

import (
	"github.com/paulmach/osm"
	"github.com/xiaokangwang/osmRoute/mapindex"
	"os"
	"sync"
)

type MapCtx struct {
	mapindex.Map
	mapFile     *os.File
	mapFileLock *sync.Mutex
}

func (c MapCtx) GetNodeWithInterconnection(Lat, Lon float64, spec ConnectionSpec) []Node {
	var nodes []Node
	_, feaIDs := c.ScanRegion(Lat, Lon, 4)
	for _, v := range feaIDs {
		node := (*c.ResolveInfoFromID(v.String())).(*osm.Node)
		absNode := c.GetNodeFromOSMNode(node)
		conn := absNode.FindConnection(spec)
		if conn != nil && len(conn) >= 1 {
			nodes = append(nodes, absNode)
		}
	}
	return nodes
}

func (c MapCtx) ResolveInfoFromID(FeaID string) *osm.Object {
	c.mapFileLock.Lock()
	Res := c.GetFeatureByID(FeaID, c.mapFile)
	c.mapFileLock.Unlock()
	return Res
}

func (c MapCtx) ListRoutes(FeaID string, spec ConnectionSpec) []Connection {
	var ret []Connection

	fromNode := (*c.ResolveInfoFromID(FeaID)).(*osm.Node)

	relations := c.GetRelationByFeature(FeaID)
	for _, v := range relations {
		info := c.ResolveInfoFromID(v.String())
		switch v.Type() {
		case osm.TypeWay:
			infoway := (*info).(*osm.Way)
			roadtype := infoway.Tags.Find("highway")

			CanCarUse := false
			CanPedestriansUse := false

			switch roadtype {
			case "motorway":
				fallthrough
			case "trunk":
				fallthrough
			case "primary":
				fallthrough
			case "motorway_link":
				fallthrough
			case "trunk_link":
				fallthrough
			case "primary_link":
				// These roads are car accessible and not accessible to pedestrians by default
				CanCarUse = true
				CanPedestriansUse = false
				break
			case "secondary":
				fallthrough
			case "tertiary":
				fallthrough
			case "residential":
				fallthrough
			case "secondary_link":
				fallthrough
			case "tertiary_link":
				fallthrough
			case "living_street":
				fallthrough
			case "road":
				fallthrough
			case "unclassified":
				// Accessible to both by default
				CanCarUse = true
				CanPedestriansUse = true
				break
			case "pedestrian":
				fallthrough
			case "footway":
				fallthrough
			case "steps":
				// Accessible to pedestrians by default
				CanCarUse = false
				CanPedestriansUse = true
				break
			case "":
				fallthrough
			default:
				continue
			}
			wayFrom := (*c.ResolveInfoFromID(infoway.Nodes[0].FeatureID().String())).(*osm.Node)
			wayTo := (*c.ResolveInfoFromID(infoway.Nodes[len(infoway.Nodes)-1].FeatureID().String())).(*osm.Node)

			if (CanCarUse && spec.CanDrive()) || (CanPedestriansUse && spec.CanWalk()) {
				ret = append(ret, c.NewConnection(fromNode, wayFrom))
				ret = append(ret, c.NewConnection(fromNode, wayTo))
			}

		case osm.TypeRelation:
			inforela := (*info).(*osm.Relation)
			_ = inforela
		}
	}
	return ret
}

func (c *MapCtx) GetNodeFromOSMNode(osmNode *osm.Node) Node {
	return &NodeImpl{
		Node: osmNode,
		c:    c,
	}
}

type NodeImpl struct {
	*osm.Node
	c *MapCtx
}

func (n NodeImpl) FindConnection(spec ConnectionSpec) []Connection {
	return n.c.ListRoutes(n.FeatureID().String(), spec)
}

type ConnectionImpl struct {
	from Node
	to   Node
}

func (c ConnectionImpl) From() Node {
	return c.from
}

func (c ConnectionImpl) To() Node {
	return c.to
}

func (c ConnectionImpl) GetCost() float64 {
	panic("implement me")
}

func (c *MapCtx) NewConnection(from, to *osm.Node) Connection {
	return ConnectionImpl{
		from: c.GetNodeFromOSMNode(from),
		to:   c.GetNodeFromOSMNode(to),
	}
}
