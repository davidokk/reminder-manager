package main

import (
	"context"
	"log"
	"net"

	"gitlab.ozon.dev/davidokk/reminder-manager/config"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	apiPkg "gitlab.ozon.dev/davidokk/reminder-manager/internal/api"

	"net/http"

	pb "gitlab.ozon.dev/davidokk/reminder-manager/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func runGRPCServer() {
	listener, err := net.Listen(config.App.GRPC.Network, config.App.GRPC.Address)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAdminServer(grpcServer, apiPkg.New())

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func runREST() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterAdminHandlerFromEndpoint(ctx, mux, config.App.REST.Endpoint, opts); err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(config.App.REST.Address, mux); err != nil {
		log.Fatal(err)
	}
}
