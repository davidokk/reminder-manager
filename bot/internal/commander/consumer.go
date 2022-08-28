package commander

import (
	"github.com/Shopify/sarama"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.ozon.dev/davidokk/reminder-manager/bot/config"
)

type consumer struct {
	bot *tgbotapi.BotAPI
}

func (c *consumer) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *consumer) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		_, _ = c.bot.Send(tgbotapi.NewMessage(config.App.Kafka.ConsumerEndpoint, string(msg.Value)))
	}
	return nil
}
