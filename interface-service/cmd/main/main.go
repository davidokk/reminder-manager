package main

import (
	"log"

	"gitlab.ozon.dev/davidokk/reminder-manager/interface-service/cmd/app"
	"gitlab.ozon.dev/davidokk/reminder-manager/interface-service/config"
	pb "gitlab.ozon.dev/davidokk/reminder-manager/interface-service/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config.ReadConfigs()

	connection, err := grpc.Dial(config.App.DataService.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewDataClient(connection)

	go app.RunREST()
	app.RunGRPCServer(client)
}
