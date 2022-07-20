package commander

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"

	"reminder-manager/config"
)

type Handler func(string) string

type Commander struct {
	bot      *tgbotapi.BotAPI
	handlers map[string]Handler
}

func Init() (*Commander, error) {
	bot, err := tgbotapi.NewBotAPI(config.ApiKey)
	if err != nil {
		return nil, errors.Wrap(err, "init tgbot")
	}

	//bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Commander{
		bot:      bot,
		handlers: make(map[string]Handler),
	}, nil
}

const UnknownCommandResponse = "Unknown command. Type /help to see the list"
const DefaultResponse = "I can only follow commands. Type /help to see the list"

func (cmd *Commander) Run() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := cmd.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if cmdName := update.Message.Command(); cmdName != "" {
				if handler, ok := cmd.handlers[cmdName]; ok {
					msg.Text = handler(update.Message.CommandArguments())
				} else {
					msg.Text = UnknownCommandResponse
				}
			} else {
				msg.Text = DefaultResponse
			}

			_, err := cmd.bot.Send(msg)
			if err != nil {
				return errors.Wrap(err, "send tg message")
			}
		}
	}

	return nil
}

func (cmd *Commander) RegisterHandler(name string, handler Handler) error {
	if _, ok := cmd.handlers[name]; ok {
		return errors.New("add handler with existing name")
	}
	cmd.handlers[name] = handler
	return nil
}
