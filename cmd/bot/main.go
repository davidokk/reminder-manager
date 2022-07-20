package main

import (
	"log"
	"reminder-manager/internal/commander"
	"reminder-manager/internal/handlers"
)

func main() {
	log.Println("start main")

	cmd, err := commander.Init()
	if err != nil {
		log.Panic(err)
	}

	handlers.AddHandlers(cmd)

	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}

}
