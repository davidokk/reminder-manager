package main

import (
	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/cmd/app"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/config"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/internal/consumer"
)

func main() {
	config.ReadConfigs()

	storage := app.ConnectRepository()
	defer storage.Close()

	go consumer.Run(sarama.NewConfig(), storage)
	app.RunGRPCServer(storage)
}
