package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"reminder-manager/config"
	"reminder-manager/internal/handlers"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.ApiKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			if cmdName := update.Message.Command(); cmdName != "" {
				if handler, ok := handlers.List[cmdName]; ok {
					msg.Text = handler(update.Message.CommandArguments())
				} else {
					msg.Text = "Unknown command. Type /help to see the list"
				}
			} else {
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
				msg.Text = "I can only follow commands. Type /help to see the list"
			}
			bot.Send(msg)
		}
	}
}
