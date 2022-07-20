package main

import (
	"log"
	"reminder-manager/internal/commander"
)

func main() {
	cmd, err := commander.Init()
	if err != nil {
		log.Panic(err)
	}

	if err := cmd.Run(); err != nil {
		log.Panic(err)
	}

}
