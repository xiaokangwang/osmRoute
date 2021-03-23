package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path"

	"github.com/xiaokangwang/osmRoute/admcommon"

	"github.com/xiaokangwang/osmRoute/util"

	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	_ "github.com/jnewmano/grpc-json-proxy/codec"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/xiaokangwang/osmRoute/mapctx"
	"github.com/xiaokangwang/osmRoute/rpc"
	"google.golang.org/grpc"
)

func logInit() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)
	grpc.EnableTracing = true
}

func grpcInit(logger *log.Entry, routeService rpc.RouteServiceServer) {
	grpc_logrus.ReplaceGrpcLogger(logger)
	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_logrus.UnaryServerInterceptor(logger),
			grpc_prometheus.UnaryServerInterceptor,
		),
		grpc_middleware.WithStreamServerChain(
			grpc_logrus.StreamServerInterceptor(logger),
			grpc_prometheus.StreamServerInterceptor,
		),
	)
	rpc.RegisterRouteServiceServer(grpcServer, routeService)

	options := []grpcweb.Option{
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true
		}),
		grpcweb.WithWebsockets(true),
		grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool {
			return true
		}),
		grpcweb.WithAllowedRequestHeaders([]string{"Wsauthtoken",
			"X-Forwarded-For", "Host", "host", "Origin", "origin"}),
	}

	wrappedGrpc := grpcweb.WrapServer(grpcServer, options...)

	router := chi.NewRouter()
	grpcWebMiddleware := util.NewGrpcWebMiddleware(wrappedGrpc)
	router.Use(
		chiMiddleware.Logger,
		chiMiddleware.Recoverer,
		grpcWebMiddleware.Handler,
	)
	router.Post("/rpc", grpcWebMiddleware.DefaultFailureHandler)

	httpServer := http.Server{
		Addr:    fmt.Sprintf("localhost:%d", 9000),
		Handler: router,
	}

	go func() {
		lis, err := net.Listen("tcp", "localhost:9001")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("failed starting http server: %v", err)
	}
}

func main() {
	logInit()
	logger := log.WithField("module", "gRPC")

	mapinde := admcommon.GetMapFromDir(path.Join(util.GetBaseDirFromEnvironment(), "testdb"))
	mapfile, err := os.Open(path.Join(util.GetBaseDirFromEnvironment(), "ireland.osm.pbf"))
	if err != nil {
		panic(err)
	}
	mapCtx := mapctx.NewMapCtx(*mapinde, mapfile)

	grpcInit(logger, &RouteService{mapctx: mapCtx, logger: log.WithField("module", "services")})
}
