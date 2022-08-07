package app

import (
	"context"
	"log"
	"net"

	"gitlab.ozon.dev/davidokk/reminder-manager/internal/storage"

	"gitlab.ozon.dev/davidokk/reminder-manager/config"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	apiPkg "gitlab.ozon.dev/davidokk/reminder-manager/internal/api"

	"net/http"

	pb "gitlab.ozon.dev/davidokk/reminder-manager/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// RunGRPCServer starts gRPC server
func RunGRPCServer(storage storage.RemindersStorage) {
	listener, err := net.Listen(config.App.GRPC.Network, config.App.GRPC.Address)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAdminServer(grpcServer, apiPkg.New(storage))

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
	if err := pb.RegisterAdminHandlerFromEndpoint(ctx, mux, config.App.HTTP.Endpoint, opts); err != nil {
		log.Fatal(err)
	}

	hmux := http.NewServeMux()
	hmux.Handle("/", mux)

	fs := http.FileServer(http.Dir("./swagger"))
	hmux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	log.Println("start listening HTTP on", config.App.HTTP.Address)
	if err := http.ListenAndServe(config.App.HTTP.Address, hmux); err != nil {
		log.Fatal(err)
	}
}
