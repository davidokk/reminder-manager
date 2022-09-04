package main

import (
	"log"

	"gitlab.ozon.dev/davidokk/reminder-manager/bot/config"
	"gitlab.ozon.dev/davidokk/reminder-manager/bot/internal/commander"
	pb "gitlab.ozon.dev/davidokk/reminder-manager/bot/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// RunBot starts the telegram bot
func RunBot(client pb.InterfaceClient) {
	cmd, err := commander.Init(config.App.Bot.APIkey, client)
	if err != nil {
		log.Fatal(err)
	}

	go cmd.RunConsumer()

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	config.ReadConfigs()

	connection, err := grpc.Dial(config.App.InterfaceService.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewInterfaceClient(connection)

	RunBot(client)
}
