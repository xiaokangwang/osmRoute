package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/xiaokangwang/osmRoute/admcommon"
	"github.com/xiaokangwang/osmRoute/mapctx"
	"github.com/xiaokangwang/osmRoute/rpc"
	"github.com/xiaokangwang/osmRoute/util"
	"google.golang.org/grpc"
)

var gRPClient rpc.RouteServiceClient

func TestRouteSpecs(t *testing.T) {
	opt := grpc.WaitForReady(true)
	rpl, err := gRPClient.Route(context.Background(), &rpc.RoutingDecisionReq{
		From: &rpc.RoutingDecisionReqLocation{Lat: 53.35214, Lon: -6.25866},
		To:   &rpc.RoutingDecisionReqLocation{Lat: 53.36135, Lon: -6.23813},
		AdditionalInfo: map[string]string{
			"objective_time":        "0.33",
			"objective_cost":        "0.33",
			"objective_sustainable": "0.33",
			"can_walkLong":          "true",
			"can_drive":             "true",
			"can_bike":              "true",
			"can_publictrans":       "true",
		},
	}, opt)
	if err != nil {
		t.Errorf("gRPClient return an error: %s", err)
	}
	if len(rpl.GetHops()) == 0 {
		t.Errorf("No hops returned")
	}
}

func TestRouteReturnAttrib(t *testing.T) {
	opt := grpc.WaitForReady(true)
	rpl, err := gRPClient.Route(context.Background(), &rpc.RoutingDecisionReq{
		From: &rpc.RoutingDecisionReqLocation{Lat: 53.35214, Lon: -6.25866},
		To:   &rpc.RoutingDecisionReqLocation{Lat: 53.36135, Lon: -6.23813},
		AdditionalInfo: map[string]string{
			"objective_time":        "0.33",
			"objective_cost":        "0.33",
			"objective_sustainable": "0.33",
			"can_walkLong":          "true",
			"can_drive":             "true",
			"can_bike":              "true",
			"can_publictrans":       "true",
		},
	}, opt)
	if err != nil {
		t.Errorf("gRPClient return an error: %s", err)
	}
	if len(rpl.GetHops()) == 0 {
		t.Errorf("No hops returned")
	}
	var found bool
	for _, v := range rpl.GetHops() {
		if v.AssociatedData != nil {
			found = true
			if methodValue, ok := v.AssociatedData["method"]; ok {
				switch methodValue {
				case "car":
					fallthrough
				case "public":
					fallthrough
				case "bus":
					fallthrough
				case "bike":
					break
				default:
					t.Error("a unrecognized value is set to method")
				}
			}
			if methodValue, ok := v.AssociatedData["co2_footprint"]; ok {
				if _, err := strconv.ParseFloat(methodValue, 64); err != nil {
					t.Error(err)
				}
			}
			if methodValue, ok := v.AssociatedData["time_saved_for_humanity"]; ok {
				if _, err := strconv.ParseFloat(methodValue, 64); err != nil {
					t.Error(err)
				}
			}
		}
	}
	assert.True(t, found, "Should have at least one route with associated data.")
}

func TestMain(m *testing.M) {
	loadEnvTest()
	logInit()
	logger := log.WithField("module", "gRPC")
	go func() {
		mapinde := admcommon.GetMapFromDir(path.Join(util.GetBaseDirFromEnvironment(), "testdb"))
		mapfile, err := os.Open(path.Join(util.GetBaseDirFromEnvironment(), "ireland.osm.pbf"))
		if err != nil {
			panic(err)
		}
		grpcInit(logger, &RouteService{mapctx: mapctx.NewMapCtx(*mapinde, mapfile), logger: log.WithField("module", "services")})
	}()

	gRPClient = rpc.NewRouteServiceClient(gRPClientSetup("localhost:9001"))
	code := m.Run()
	os.Exit(code)
}

func gRPClientSetup(host string) *grpc.ClientConn {
	opt := []grpc.DialOption{grpc.WithInsecure()}
	cc, err := grpc.Dial(host, opt...)
	if err != nil {
		panic(err)
	}
	return cc
}

func loadEnvTest() {
	cwd, _ := os.Getwd()
	errEnv := godotenv.Load(string(cwd) + `/.env`)
	if errEnv != nil {
		log.Fatalf("Failed to load environment .env")
		panic(errEnv)
	}
}
