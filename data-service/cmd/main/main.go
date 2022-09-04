package main

import (
	"log"
	"net/http"

	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/internal/storage/cached"

	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/cmd/app"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/config"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/internal/consumer"

	_ "expvar"
)

func main() {
	config.ReadConfigs()

	repo := app.ConnectRepository()
	defer repo.Close()

	storage := cached.New(repo)

	go func() {
		err := http.ListenAndServe(config.App.PprofAddress, nil)
		if err != nil {
			log.Println(err.Error())
		}
	}()

	go consumer.Run(sarama.NewConfig(), storage)
	app.RunGRPCServer(storage)
}
