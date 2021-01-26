package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/xiaokangwang/osmRoute/util"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/paulmach/osm"
	"github.com/xiaokangwang/osmRoute/adm"
	"github.com/xiaokangwang/osmRoute/mapctx"
	"github.com/xiaokangwang/osmRoute/rpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	grpcServer := grpc.NewServer()

	mapinde := adm.GetMapFromDir(util.GetBaseDirFromEnvironment() + "/testdb")
	mapfile, err := os.Open(util.GetBaseDirFromEnvironment() + "/ireland.osm.pbf")
	if err != nil {
		panic(err)
	}
	mapCtx := mapctx.NewMapCtx(*mapinde, mapfile)

	rpc.RegisterRouteServiceServer(grpcServer, &RouteService{mapctx: mapCtx})

	wrappedGrpc := grpcweb.WrapServer(grpcServer, grpcweb.WithWebsockets(true), grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool {
		return true
	}))

	handler := func(resp http.ResponseWriter, req *http.Request) {
		req.Header.Set("Upgrade", strings.ToLower(req.Header.Get("Upgrade")))
		log.Println(fmt.Sprintf("%q %q %q %q", req.Host, req.Method, req.Proto, req.URL))
		wrappedGrpc.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", 9000),
		Handler: http.HandlerFunc(handler),
	}

	go func() {
		lis, err := net.Listen("tcp", ":9001")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()

	if err := httpServer.ListenAndServe(); err != nil {
		grpclog.Fatalf("failed starting http server: %v", err)
	}
}

type RouteService struct {
	rpc.UnimplementedRouteServiceServer
	mapctx *mapctx.MapCtx
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
		by, err := json.Marshal(objWay)
		if err != nil {
			panic(err)
		}
		println(string(by))
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
		println(string(by))
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
	panic("implement me")
}

func (r RouteService) SearchByNameExact(ctx context.Context, search *rpc.NameSearch) (*rpc.ObjectList, error) {
	panic("implement me")
}
