package main

import (
	"log"
	"net"
	"os"

	apiPkg "gitlab.ozon.dev/davidokk/reminder-manager/internal/api"
	pb "gitlab.ozon.dev/davidokk/reminder-manager/pkg/api"
	"google.golang.org/grpc"

	"gitlab.ozon.dev/davidokk/reminder-manager/internal/commander"
	"gitlab.ozon.dev/davidokk/reminder-manager/internal/handlers"
)

func runBot() {
	cmd, err := commander.Init(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	handlers.AddHandlers(cmd)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

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

func main() {
	runGRPCServer()
}
