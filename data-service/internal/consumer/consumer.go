package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/config"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/internal/producer"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/internal/storage"
)

// Run starts consumer reads messages from data incoming topic
func Run(cfg *sarama.Config, storage storage.RemindersStorage) {
	ctx := context.Background()

	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest

	client, err := sarama.NewConsumerGroup(config.App.Kafka.Brokers, config.App.Kafka.ConsumerGroupID, cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}

	consumer := &consumer{
		storage:  storage,
		producer: producer.New(cfg),
	}

	for {
		if err := client.Consume(ctx, []string{config.App.Kafka.DataIncomingTopic}, consumer); err != nil {
			log.Printf("on consume: %v", err)
			time.Sleep(time.Second * 2)
		}
	}
}

type consumer struct {
	storage  storage.RemindersStorage
	producer sarama.AsyncProducer
}

func (c *consumer) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *consumer) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		params := make(map[string]string)
		if err := json.Unmarshal(msg.Value, &params); err != nil {
			log.Printf("unmarshal error: %s\n", err.Error())
		}

		ctx := context.Background()
		response := "success"

		switch string(msg.Key) {
		case "update":
			if err := update(ctx, c.storage, params); err != nil {
				response = fmt.Sprintf("error: %s", err.Error())
			}
		case "remove":
			if err := remove(ctx, c.storage, params); err != nil {
				response = fmt.Sprintf("error: %s", err.Error())
			}
		case "create":
			if ID, err := create(ctx, c.storage, params); err != nil {
				response = fmt.Sprintf("error: %s", err.Error())
			} else {
				response = fmt.Sprintf("succes, id = %d", ID)
			}
		}

		c.producer.Input() <- &sarama.ProducerMessage{
			Topic: config.App.Kafka.DataResponseTopic,
			Key:   sarama.ByteEncoder(msg.Key),
			Value: sarama.StringEncoder(response),
		}
	}
	return nil
}
