package app

import (
	"log"
	"net"

	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/config"
	apiPkg "gitlab.ozon.dev/davidokk/reminder-manager/data-service/internal/api"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/internal/storage"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/pkg/api"

	"google.golang.org/grpc"
)

// RunGRPCServer starts gRPC server
func RunGRPCServer(storage storage.RemindersStorage) {
	listener, err := net.Listen(config.App.GRPC.Network, config.App.GRPC.Address)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterDataServer(grpcServer, apiPkg.New(storage))

	log.Println("start listening gRPC on", config.App.GRPC.Address)
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
