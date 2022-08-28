//go:build integration
// +build integration

package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var App struct {
	Address  string `yaml:"address"`
	Postgres struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"db-name"`
	} `yaml:"postgres"`
}

// ReadConfigs gets app's configs
func ReadConfigs() {
	file, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := yaml.Unmarshal(file, &App); err != nil {
		log.Fatal(err.Error())
	}
	log.Println("configs read")
}
