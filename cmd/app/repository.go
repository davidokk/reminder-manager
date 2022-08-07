package app

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/davidokk/reminder-manager/config"
	"gitlab.ozon.dev/davidokk/reminder-manager/internal/storage/repository/postgres"
)

// ConnectRepository connects to db
func ConnectRepository() *postgres.Repository {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.App.Postgres.Host, config.App.Postgres.Port,
		config.App.Postgres.User, config.App.Postgres.Password,
		config.App.Postgres.DBName,
	)

	pool, err := pgxpool.Connect(ctx, psqlConn)
	if err != nil {
		log.Fatal("can't connect to database ", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("ping database error ", err)
	} else {
		log.Println("connected to postgres")
	}

	return postgres.NewRepository(pool)
}
