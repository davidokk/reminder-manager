package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// App is a group of app's configs
var App struct {
	Storage struct {
		WaitingTime time.Duration `yaml:"waiting-time"`
		PoolSize    uint8         `yaml:"pool-size"`
	} `yaml:"storage"`
	GRPC struct {
		Network string `yaml:"network"`
		Address string `yaml:"address"`
	} `yaml:"grpc"`
	Postgres struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"db-name"`
	} `yaml:"postgres"`
	Kafka struct {
		DataIncomingTopic string   `yaml:"data-incoming-topic"`
		DataResponseTopic string   `yaml:"data-response-topic"`
		ConsumerGroupID   string   `yaml:"consumer-group-id"`
		Brokers           []string `yaml:"brokers"`
	} `yaml:"kafka"`
}

// ReadConfigs gets app's configs
func ReadConfigs() {
	file, err := os.ReadFile("data-service/config/config.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := yaml.Unmarshal(file, &App); err != nil {
		log.Fatal(err.Error())
	}
	log.Println("configs read")
}
