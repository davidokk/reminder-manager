package config

import (
	"log"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// App is a group of app's configs
var App struct {
	Bot struct {
		APIKey string
	}
	Storage struct {
		WaitingTime time.Duration `yaml:"waiting-time"`
		PoolSize    uint8         `yaml:"pool-size"`
	}
	GRPC struct {
		Network string
		Address string
	}
	REST struct {
		Endpoint string
		Address  string
	}
}

var required = []string{
	"bot:apikey",

	"storage:waiting-time",
	"storage:pool-size",

	"grpc:network",
	"grpc:address",

	"rest:endpoint",
	"rest:address",
}

func checkValid(date []byte) {
	config := make(map[string]map[string]string)
	if err := yaml.Unmarshal(date, &config); err != nil {
		log.Fatal(err.Error())
	}
	for _, str := range required {
		p := strings.Split(str, ":")
		if _, ok := config[p[0]][p[1]]; !ok {
			log.Fatalf("%s field not specified", str)
		}
	}
}

// ReadConfigs gets app configs
func ReadConfigs() {
	file, err := os.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}
	checkValid(file)
	if err := yaml.Unmarshal(file, &App); err != nil {
		log.Fatal(err.Error())
	}
	log.Println("configs read")
}
