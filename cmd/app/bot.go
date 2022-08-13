package app

import (
	"log"

	"gitlab.ozon.dev/davidokk/reminder-manager/config"
	"gitlab.ozon.dev/davidokk/reminder-manager/internal/commander"
	storagePkg "gitlab.ozon.dev/davidokk/reminder-manager/internal/storage"
)

// RunBot starts the telegram bot
func RunBot(storage storagePkg.RemindersStorage) {
	cmd, err := commander.Init(config.App.Bot.APIKey, storage)
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
