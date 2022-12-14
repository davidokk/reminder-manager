package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// App is a group of app's configs
var App struct {
	Bot struct {
		APIkey string `yaml:"apikey"`
	}
	InterfaceService struct {
		Address string `yaml:"address"`
	} `yaml:"interface-service"`
	Kafka struct {
		DataResponseTopic string   `yaml:"data-response-topic"`
		ConsumerGroupID   string   `yaml:"consumer-group-id"`
		Brokers           []string `yaml:"brokers"`
		ConsumerEndpoint  int64    `yaml:"consumer-endpoint"`
	} `yaml:"kafka"`
}

// ReadConfigs gets app's configs
func ReadConfigs() {
	file, err := os.ReadFile("bot/config/config.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := yaml.Unmarshal(file, &App); err != nil {
		log.Fatal(err.Error())
	}
	log.Println("configs read")
}
