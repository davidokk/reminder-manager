package main

import (
	"log"
	"os"

	"gitlab.ozon.dev/davidokk/reminder-manager/internal/commander"
	"gitlab.ozon.dev/davidokk/reminder-manager/internal/handlers"
)

func main() {
	log.Println("start main")

	cmd, err := commander.Init(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	handlers.AddHandlers(cmd)

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
