package app

import (
	"context"
	"log"
	"net"

	"gitlab.ozon.dev/davidokk/reminder-manager/interface-service/config"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	apiPkg "gitlab.ozon.dev/davidokk/reminder-manager/interface-service/internal/api"

	"net/http"

	pb "gitlab.ozon.dev/davidokk/reminder-manager/interface-service/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// RunGRPCServer starts gRPC server
func RunGRPCServer(client pb.DataClient) {
	listener, err := net.Listen(config.App.GRPC.Network, config.App.GRPC.Address)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterInterfaceServer(grpcServer, apiPkg.New(client))

	log.Println("start listening gRPC on", config.App.GRPC.Address)
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

// RunREST starts HTTP
func RunREST() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterInterfaceHandlerFromEndpoint(ctx, mux, config.App.HTTP.Endpoint, opts); err != nil {
		log.Fatal(err)
	}

	hmux := http.NewServeMux()
	hmux.Handle("/", mux)

	fs := http.FileServer(http.Dir("interface-service/swagger"))
	hmux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	log.Println("start listening HTTP on", config.App.HTTP.Address)
	if err := http.ListenAndServe(config.App.HTTP.Address, hmux); err != nil {
		log.Fatal(err)
	}
}
