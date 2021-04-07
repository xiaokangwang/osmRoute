package main

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/beefsack/go-astar"
	"github.com/paulmach/osm"
	log "github.com/sirupsen/logrus"
	"github.com/xiaokangwang/osmRoute/attributes"
	"github.com/xiaokangwang/osmRoute/mapctx"
	"github.com/xiaokangwang/osmRoute/rpc"
)

type RouteService struct {
	rpc.UnimplementedRouteServiceServer
	mapctx *mapctx.MapCtx
	logger *log.Entry
}

func (r RouteService) Route(ctx context.Context, req *rpc.RoutingDecisionReq) (*rpc.RoutingDecisionResp, error) {
	var ret rpc.RoutingDecisionResp
	var additionalInfo = req.GetAdditionalInfo()
	spec, err := attributes.ParseRoutingInputAttribute(additionalInfo)
	if err != nil {
		ret.Code = -1
		ret.Msg = fmt.Sprintf("%x", err)
		return &ret, nil
	}

	r.mapctx.SetSpec(spec)

	InitialNodes := r.mapctx.GetNodeWithInterconnection(req.From.Lat, req.From.Lon, spec)

	for _, v := range InitialNodes {
		node := v.(*mapctx.NodeImpl)
		r.logger.Debugln(node.FeatureID().String())
	}

	if len(InitialNodes) == 0 {
		ret.Code = -2
		ret.Msg = "InitialNodes length == 0"
		return &ret, nil
	}

	InitialNodesF := r.mapctx.GetNodeWithInterconnection(req.To.Lat, req.To.Lon, spec)

	for _, v := range InitialNodesF {
		node := v.(*mapctx.NodeImpl)
		r.logger.Debugln(node.FeatureID().String())
	}

	if len(InitialNodesF) == 0 {
		ret.Code = -3
		ret.Msg = "InitialNodesF length == 0"
		return &ret, nil
	}

	path, dist, found := astar.Path(InitialNodes[0], InitialNodesF[0])
	r.logger.Debugln(found)
	r.logger.Debugln(dist)
	var last astar.Pather
	reverseAny(path)
	for _, v := range path {

		var hop rpc.RoutingDecision

		if last != nil {
			fid := ""
			ViaObject := last.(mapctx.Node).PathNeighborVia(v)
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
			r.logger.Debugln("via:", fid)
			hop.Via = fid
			hop.AssociatedData = last.(*mapctx.NodeImpl).PathNeighborConnection(v).GetAttributes()
		}
		r.logger.Debugln(v.(*mapctx.NodeImpl).FeatureID().String())
		last = v
		hop.From = v.(*mapctx.NodeImpl).FeatureID().String()

		ret.Hops = append(ret.Hops, &hop)
	}
	ret.Code = 0
	ret.Msg = "Success"
	return &ret, nil
}

func reverseAny(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

type KnownPoint struct {
	Lon float64
	Lat float64
}

func (r RouteService) Resolve(ctx context.Context, request *rpc.ObjectResolveRequest) (*rpc.ReturnedObject, error) {
	obj := r.mapctx.ResolveInfoFromID(request.FeatureID)

	if obj == nil {
		return &rpc.ReturnedObject{
			FeatureID:     "",
			ObjectContent: nil,
			Found:         false,
		}, nil
	}
	switch (*obj).ObjectID().Type() {
	case osm.TypeWay:
		objWay := (*obj).(*osm.Way)
		//Sample some way
		var knownPoints []KnownPoint
		length := len(objWay.Nodes)
		resolve := func(i int) {
			osmobj := (*r.mapctx.ResolveInfoFromID(objWay.Nodes[i].FeatureID().String())).(*osm.Node)
			objWay.Nodes[i].Lon = osmobj.Lon
			objWay.Nodes[i].Lat = osmobj.Lat
			knownPoints = append(knownPoints, KnownPoint{
				Lon: osmobj.Lon,
				Lat: osmobj.Lat,
			})
		}
		if length <= 2 {
			for i, _ := range objWay.Nodes {
				resolve(i)
			}
		} else {
			resolve(0)
			resolve(length / 2)
			resolve(length - 1)
		}

		byw, err := json.Marshal(knownPoints)
		if err != nil {
			panic(err)
		}
		objWay.Tags = append(objWay.Tags, osm.Tag{
			Key:   "X-osmRoute-KnownPoints",
			Value: string(byw),
		})
		by, err := json.Marshal(objWay)
		if err != nil {
			panic(err)
		}
		r.logger.Debugln(string(by))
		return &rpc.ReturnedObject{
			FeatureID:     request.FeatureID,
			ObjectContent: by,
			Found:         true,
		}, nil
	case osm.TypeNode:
		objNode := (*obj).(*osm.Node)
		by, err := json.Marshal(objNode)
		if err != nil {
			panic(err)
		}
		r.logger.Debugln(string(by))
		return &rpc.ReturnedObject{
			FeatureID:     request.FeatureID,
			ObjectContent: by,
			Found:         true,
		}, nil
	case osm.TypeRelation:
		objRelation := (*obj).(*osm.Relation)
		by, err := json.Marshal(objRelation)
		if err != nil {
			panic(err)
		}
		return &rpc.ReturnedObject{
			FeatureID:     request.FeatureID,
			ObjectContent: by,
			Found:         true,
		}, nil
	}
	panic("Unreachable")

}

func (r RouteService) ScanRegion(ctx context.Context, request *rpc.ScanRegionRequest) (*rpc.ObjectListWithAssociatedObjects, error) {
	_, obj, loc := r.mapctx.ScanRegion(request.Lat, request.Lon, 4)
	return &rpc.ObjectListWithAssociatedObjects{FeatureID: func() []string {
		var ret []string
		for _, v := range obj {
			ret = append(ret, v.String())
		}
		return ret
	}(), FeatureIDAndAssociatedObjects: func() map[string]*rpc.ObjectList {
		var ret = make(map[string]*rpc.ObjectList)
		for _, v := range obj {
			reps, _ := r.GetAssociatedObject(ctx, &rpc.GetAssociatedObjectRequest{FeatureID: v.String()})
			ret[v.String()] = reps
		}
		return ret
	}(),
		LocationAssociation: func() []*rpc.LocationAssociation {
			var ret []*rpc.LocationAssociation
			for _, v := range loc {
				ret = append(ret, &rpc.LocationAssociation{
					Nodes: &rpc.ObjectList{FeatureID: func() []string {
						var retw []string
						for _, v2 := range v.FeatureIDs {
							retw = append(retw, v2.String())
						}
						return retw
					}()},
					Lat: v.Lat,
					Lon: v.Lon,
				})
			}
			return ret
		}(),
	}, nil
}

func (r RouteService) GetAssociatedObject(ctx context.Context, request *rpc.GetAssociatedObjectRequest) (*rpc.ObjectList, error) {
	ids := r.mapctx.GetRelationByFeature(request.FeatureID)
	return &rpc.ObjectList{FeatureID: func() []string {
		var ret []string
		for _, v := range ids {
			ret = append(ret, v.String())
		}
		return ret
	}()}, nil
}

func (r RouteService) SearchByNamePrefix(ctx context.Context, search *rpc.NameSearch) (*rpc.NameList, error) {
	keywordPrefix := search.Keyword
	results := r.mapctx.SearchByNamePrefix(keywordPrefix)
	return &rpc.NameList{ObjectName: results}, nil
}

func (r RouteService) SearchByNameExact(ctx context.Context, search *rpc.NameSearch) (*rpc.ObjectList, error) {
	keywordPrefix := search.Keyword
	results, _ := r.mapctx.SearchByName(keywordPrefix)
	for _, v := range results {
		r.logger.Debugln(v.String())
	}
	return &rpc.ObjectList{FeatureID: func() []string {
		var ret []string
		for _, v := range results {
			ret = append(ret, v.String())
		}
		return ret
	}()}, nil

}
