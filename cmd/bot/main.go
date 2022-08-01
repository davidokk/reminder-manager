package main

import (
	"log"

	"gitlab.ozon.dev/davidokk/reminder-manager/config"
	"gitlab.ozon.dev/davidokk/reminder-manager/internal/commander"
	"gitlab.ozon.dev/davidokk/reminder-manager/internal/handlers"
	"gitlab.ozon.dev/davidokk/reminder-manager/internal/storage"
)

func runBot() {
	cmd, err := commander.Init(config.App.Bot.APIKey)
	if err != nil {
		log.Fatal(err)
	}

	handlers.AddHandlers(cmd)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	config.ReadConfigs()

	storage.Init()

	go runBot()
	go runREST()
	runGRPCServer()
}
