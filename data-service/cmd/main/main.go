package main

import (
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/cmd/app"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/config"
)

func main() {
	config.ReadConfigs()

	storage := app.ConnectRepository()
	defer storage.Close()

	app.RunGRPCServer(storage)
}
