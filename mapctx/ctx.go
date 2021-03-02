package mapctx

import (
	"errors"
	"github.com/beefsack/go-astar"
	"github.com/paulmach/osm"
	"github.com/xiaokangwang/osmRoute/mapindex"
	"github.com/xiaokangwang/osmRoute/util"
	"math"
	"os"
	"sort"
	"strings"
	"sync"
)

type MapCtx struct {
	mapindex.Map
	mapFile     *os.File
	mapFileLock *sync.Mutex
	mapNode     map[string]*NodeImpl
	spec        ConnectionSpec
}

func NewMapCtx(maps mapindex.Map, mapFile *os.File) *MapCtx {
	return &MapCtx{
		Map:         maps,
		mapFile:     mapFile,
		mapFileLock: &sync.Mutex{},
		mapNode:     map[string]*NodeImpl{},
	}
}

func (c MapCtx) GetNodeWithInterconnection(Lat, Lon float64, spec ConnectionSpec) []Node {
	return c.GetNodeWithInterconnection4(Lat, Lon, spec, 3)
}

func (c MapCtx) GetNodeWithInterconnection4(Lat, Lon float64, spec ConnectionSpec, mask int) []Node {
	var feaIDs osm.FeatureIDs
	var acceptedNodes []Node
	_, _, feaG := c.ScanRegion(Lat, Lon, mask)

	sort.Sort(NodeGDistanceSlice{
		nodes:     feaG,
		OriginLat: Lat,
		OriginLon: Lon,
	})

	for _, v := range feaG {
		feaIDs = append(feaIDs, v.FeatureIDs...)
	}

	for _, absNode := range feaIDs {
		node := (*c.ResolveInfoFromID(absNode.String())).(*osm.Node)
		absNode := c.GetNodeFromOSMNode(node)
		conn := absNode.FindConnection(spec)
		if conn != nil && len(conn) >= 1 {
			acceptedNodes = append(acceptedNodes, absNode)
		}
		if len(acceptedNodes) >= 3 {
			break
		}
	}

	return acceptedNodes
}

func (c *MapCtx) SetSpec(spec ConnectionSpec) {
	c.spec = spec
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
	ret = c.CreateInterconnections(relations, spec, ret, fromNode, []osm.FeatureID{})

	//Scan if there is a bus station

	_, _, ScanedResult := c.ScanRegion(fromNode.Lat, fromNode.Lon, 3)

	for _, v := range ScanedResult {
		for _, r := range v.SignificantRelations {
			ret = c.CreateInterconnections([]osm.FeatureID{r}, spec, ret, fromNode, v.FeatureIDs)
		}

	}

	//For Bus stop, we allow user to wrap to a nearby location
	if spec.CanPublicTransport() && fromNode.Tags.Find("bus") == "yes" {
		nodes := c.GetNodeWithInterconnection4(fromNode.Lat, fromNode.Lon, specProxyNoBus{spec}, 2)

		if len(nodes) > 0 {
			for _, node := range nodes[:1] {
				ret = append(ret, c.ListRoutes(node.(*NodeImpl).FeatureID().String(), specProxyNoBus{spec})...)
			}
		}
	}

	return ret
}

type specProxyNoBus struct {
	ConnectionSpec
}

func (s specProxyNoBus) CanPublicTransport() bool {
	return false
}

func (c MapCtx) CreateInterconnections(relations osm.FeatureIDs, spec ConnectionSpec, ret []Connection, fromNode *osm.Node, associatedNodeID osm.FeatureIDs) []Connection {
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
				ret = append(ret, c.NewConnection(fromNode, wayFrom, *info))
				ret = append(ret, c.NewConnection(fromNode, wayTo, *info))

				//If there is an seen node, we will generate a route even if it is not the starting or ending

				for _, v := range infoway.Nodes {
					name := v.FeatureID().String()
					if _, ok := c.mapNode[name]; ok {
						SeenNode := (*c.ResolveInfoFromID(name)).(*osm.Node)
						ret = append(ret, c.NewConnection(fromNode, SeenNode, *info))
					}
				}
			}

		case osm.TypeRelation:
			inforela := (*info).(*osm.Relation)
			_ = inforela

			routevalue := inforela.Tags.Find("route")

			if routevalue == "bus" && spec.CanPublicTransport() {
				if member, errw := checkIntersectionAssociation(inforela.Members, associatedNodeID); errw == nil {
					//OK, now generate a route for all other stations
					startexact := (*c.ResolveInfoFromID(member.FeatureID().String())).(*osm.Node)
					for _, memberw := range inforela.Members {
						switch memberw.Role {
						case "platform":
							fallthrough
						case "platform_exit_only":
							ending := (*c.ResolveInfoFromID(memberw.FeatureID().String())).(*osm.Node)
							connection := c.NewConnection5(fromNode, startexact, ending, *info, 1)
							_ = connection
							ret = append(ret, connection)
						}
					}

				}
			}
		}
	}
	return ret
}

func checkIntersectionAssociation(relationMember osm.Members, candidate osm.FeatureIDs) (osm.Member, error) {
	for _, v := range relationMember {
		for _, v2 := range candidate {
			if v.FeatureID() == v2 && strings.HasPrefix(v.Role, "platform") {
				switch v.Role {
				case "platform":
					fallthrough
				case "platform_exit_only":
					return v, nil
				}
			}
		}
	}
	return osm.Member{}, errors.New("not found")
}

func (c *MapCtx) GetNodeFromOSMNode(osmNode *osm.Node) Node {
	if val, ok := c.mapNode[osmNode.FeatureID().String()]; ok {
		return val
	}
	newnode := &NodeImpl{
		Node: osmNode,
		c:    c,
	}
	c.mapNode[osmNode.FeatureID().String()] = newnode
	return newnode
}

type NodeImpl struct {
	*osm.Node
	c          *MapCtx
	cachedConn []Connection
}

func (n *NodeImpl) PathNeighbors() []astar.Pather {
	conns := n.FindConnection(n.c.spec)
	n.cachedConn = conns
	var ret []astar.Pather
	for _, v := range conns {
		ret = append(ret, v.To())
	}
	return ret
}

func (n NodeImpl) PathNeighborCost(to astar.Pather) float64 {
	for _, v := range n.cachedConn {
		if v.To() == to {
			return v.GetCost()
		}
	}
	return math.Inf(1)
}

func (n NodeImpl) PathNeighborVia(to astar.Pather) osm.Object {
	for _, v := range n.cachedConn {
		if v.To() == to {
			return v.Via()
		}
	}
	return nil
}

func (n NodeImpl) PathEstimatedCost(to astar.Pather) float64 {
	return util.GPStoMeter(n.Lat, n.Lon, to.(*NodeImpl).Lat, to.(*NodeImpl).Lon)
}

func (n NodeImpl) FindConnection(spec ConnectionSpec) []Connection {
	return n.c.ListRoutes(n.FeatureID().String(), spec)
}

type ConnectionImpl struct {
	from         Node
	to           Node
	via          osm.Object
	fromExact    Node
	costDiscount float64
}

func (c ConnectionImpl) From() Node {
	return c.from
}

func (c ConnectionImpl) To() Node {
	return c.to
}

func (c ConnectionImpl) Via() osm.Object {
	return c.via
}

func (c ConnectionImpl) GetCost() float64 {

	return util.GPStoMeter(c.from.(*NodeImpl).Lat, c.from.(*NodeImpl).Lon, c.to.(*NodeImpl).Lat, c.to.(*NodeImpl).Lon) * c.costDiscount
}

func (c *MapCtx) NewConnection(from, to *osm.Node, via osm.Object) Connection {
	return ConnectionImpl{
		from:         c.GetNodeFromOSMNode(from),
		to:           c.GetNodeFromOSMNode(to),
		via:          via,
		costDiscount: 1,
	}
}

func (c *MapCtx) NewConnection5(from, to, exact *osm.Node, via osm.Object, discount float64) Connection {
	return ConnectionImpl{
		from:         c.GetNodeFromOSMNode(from),
		to:           c.GetNodeFromOSMNode(to),
		via:          via,
		fromExact:    c.GetNodeFromOSMNode(exact),
		costDiscount: discount,
	}
}

type NodeDistanceSlice struct {
	nodes     []Node
	OriginLat float64
	OriginLon float64
}

func (p NodeDistanceSlice) Len() int { return len(p.nodes) }

func (p NodeDistanceSlice) Less(i, j int) bool {
	iLon := p.nodes[i].(*NodeImpl).Lon
	iLat := p.nodes[i].(*NodeImpl).Lat

	jLon := p.nodes[j].(*NodeImpl).Lon
	jLat := p.nodes[j].(*NodeImpl).Lat

	return util.GPStoMeter(iLon, iLat, p.OriginLon, p.OriginLat) < util.GPStoMeter(jLon, jLat, p.OriginLon, p.OriginLat)
}

func (p NodeDistanceSlice) Swap(i, j int) { p.nodes[i], p.nodes[j] = p.nodes[j], p.nodes[i] }

type NodeGDistanceSlice struct {
	nodes     []mapindex.FeatureIDWithLocation
	OriginLat float64
	OriginLon float64
}

func (p NodeGDistanceSlice) Len() int { return len(p.nodes) }

func (p NodeGDistanceSlice) Less(i, j int) bool {
	iLon := p.nodes[i].Lon
	iLat := p.nodes[i].Lat

	jLon := p.nodes[j].Lon
	jLat := p.nodes[j].Lat

	return util.GPStoMeter(iLon, iLat, p.OriginLon, p.OriginLat) < util.GPStoMeter(jLon, jLat, p.OriginLon, p.OriginLat)
}

func (p NodeGDistanceSlice) Swap(i, j int) { p.nodes[i], p.nodes[j] = p.nodes[j], p.nodes[i] }
