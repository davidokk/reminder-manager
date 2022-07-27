package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	apiPkg "gitlab.ozon.dev/davidokk/reminder-manager/internal/api"
	pb "gitlab.ozon.dev/davidokk/reminder-manager/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func runGRPCServer() {
	listener, err := net.Listen("tcp", ":8081")
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
	if err := pb.RegisterAdminHandlerFromEndpoint(ctx, mux, ":8081", opts); err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
