package producer

import (
	"log"

	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/davidokk/reminder-manager/interface-service/config"
)

func New() sarama.AsyncProducer {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true

	asyncProducer, err := sarama.NewAsyncProducer(config.App.Kafka.Brokers, cfg)
	if err != nil {
		log.Fatalf("asyn kafka: %v", err)
	}

	go func() {
		for msg := range asyncProducer.Errors() {
			log.Printf("sent error: %s", msg.Error())
		}
	}()

	go func() {
		for res := range asyncProducer.Successes() {
			log.Printf("sent success: %+v\n", res)
		}
	}()

	return asyncProducer
}
