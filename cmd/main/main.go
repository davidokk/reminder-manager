package main

import (
	"gitlab.ozon.dev/davidokk/reminder-manager/cmd/app"
	"gitlab.ozon.dev/davidokk/reminder-manager/config"
)

func main() {
	config.ReadConfigs()

	storage := app.ConnectRepository()
	defer storage.Close()

	go app.RunBot(storage)
	go app.RunREST()
	app.RunGRPCServer(storage)
}
