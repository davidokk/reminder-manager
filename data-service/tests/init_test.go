//go:build integration
// +build integration

package tests

import (
	"log"
	"time"

	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/pkg/api"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/tests/config"
	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/tests/postgres"
	"google.golang.org/grpc"
)

var Client api.DataClient
var Db *postgres.DB

func init() {
	config.ReadConfigs()

	connection, err := grpc.Dial(config.App.Address, grpc.WithInsecure(), grpc.WithTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err.Error())
	}

	Db = postgres.New()
	Client = api.NewDataClient(connection)
}
