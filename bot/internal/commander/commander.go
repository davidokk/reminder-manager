package commander

import (
	"context"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/davidokk/reminder-manager/bot/config"

	pb "gitlab.ozon.dev/davidokk/reminder-manager/bot/pkg/api"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

// Handler represents the telegram bot command handler function
type Handler func(params string, client pb.InterfaceClient) string

// Commander allows you to initialize and run the telegram bot
type Commander struct {
	bot      *tgbotapi.BotAPI
	handlers map[string]Handler
	service  pb.InterfaceClient
}

// Init connects to the bot with the given key
func Init(apiKey string, client pb.InterfaceClient) (*Commander, error) {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		return nil, errors.Wrap(err, "init tgbot")
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	cmd := &Commander{
		bot:      bot,
		handlers: make(map[string]Handler),
		service:  client,
	}

	AddHandlers(cmd)

	return cmd, nil
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
					msg.Text = handler(update.Message.CommandArguments(), cmd.service)
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

// RunConsumer stars consumer reads messages from data responses topic
func (cmd *Commander) RunConsumer() {
	ctx := context.Background()

	cfg := sarama.NewConfig()

	client, err := sarama.NewConsumerGroup(config.App.Kafka.Brokers, config.App.Kafka.ConsumerGroupID, cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}

	consumer := &consumer{
		bot: cmd.bot,
	}

	for {
		if err := client.Consume(ctx, []string{config.App.Kafka.DataResponseTopic}, consumer); err != nil {
			log.Printf("on consume: %v", err)
			time.Sleep(time.Second * 2)
		}
	}
}

// RegisterHandler adds a new Handler into Commander
func (cmd *Commander) RegisterHandler(name string, handler Handler) {
	if _, ok := cmd.handlers[name]; ok {
		log.Printf("add handler with existing name - %s\n", name)
	}
	cmd.handlers[name] = handler
}
