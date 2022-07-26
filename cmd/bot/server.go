package main

import (
	"net"

	apiPkg "gitlab.ozon.dev/davidokk/reminder-manager/internal/api"
	pb "gitlab.ozon.dev/davidokk/reminder-manager/pkg/api"
	"google.golang.org/grpc"
)

func runGRPCServer() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAdminServer(grpcServer, apiPkg.New())

	if err = grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
