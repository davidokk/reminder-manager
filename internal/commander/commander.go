package commander

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

// Handler represents the telegram bot command handler function
type Handler func(string) string

// Commander allows you to initialize and run the telegram bot
type Commander struct {
	bot      *tgbotapi.BotAPI
	handlers map[string]Handler
}

// Init connects to the bot with the given key
func Init(apiKey string) (*Commander, error) {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		return nil, errors.Wrap(err, "init tgbot")
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Commander{
		bot:      bot,
		handlers: make(map[string]Handler),
	}, nil
}

const unknownCommandResponse = "Unknown command. Type /help to see the list"
const defaultResponse = "I can only follow commands. Type /help to see the list"

const botDefaultOffset = 0
const botDefaultTimeout = 60

// Run launch telegram bot
func (cmd *Commander) Run() error {
	u := tgbotapi.NewUpdate(botDefaultOffset)
	u.Timeout = botDefaultTimeout

	updates := cmd.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if cmdName := update.Message.Command(); cmdName != "" {
				if handler, ok := cmd.handlers[cmdName]; ok {
					msg.Text = handler(update.Message.CommandArguments())
				} else {
					msg.Text = unknownCommandResponse
				}
			} else {
				msg.Text = defaultResponse
			}

			_, err := cmd.bot.Send(msg)
			if err != nil {
				return errors.Wrap(err, "send tg message")
			}
		}
	}

	return nil
}

// RegisterHandler adds a new Handler into Commander
func (cmd *Commander) RegisterHandler(name string, handler Handler) {
	if _, ok := cmd.handlers[name]; ok {
		log.Println(errors.New("add handler with existing name"))
	}
	cmd.handlers[name] = handler
}
