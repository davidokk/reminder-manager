package commander

import (
	"log"
	"reminder-manager/internal/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"

	"reminder-manager/config"
)

type Commander struct {
	bot *tgbotapi.BotAPI
}

func Init() (*Commander, error) {
	bot, err := tgbotapi.NewBotAPI(config.ApiKey)
	if err != nil {
		return nil, errors.Wrap(err, "init tgbot")
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Commander{
		bot: bot,
	}, nil
}

func (cmd *Commander) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := cmd.bot.GetUpdatesChan(u)

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
			_, err := cmd.bot.Send(msg)
			if err != nil {
				return errors.Wrap(err, "send tg message")
			}
		}
	}

	return nil
}
