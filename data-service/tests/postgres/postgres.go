//go:build integration
// +build integration

package postgres

import (
	"fmt"
	"log"
	"sync"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/net/context"

	"gitlab.ozon.dev/davidokk/reminder-manager/data-service/tests/config"
)

type DB struct {
	sync.Mutex
	pool *pgxpool.Pool
}

func New() *DB {
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
	}

	return &DB{
		pool: pool,
	}
}

func (db *DB) SetUp(t *testing.T) {
	t.Helper()
	ctx := context.Background()
	db.Lock()
	db.Truncate(ctx)
}

func (db *DB) TearDown() {
	defer db.Unlock()
	db.Truncate(context.Background())
}

func (db *DB) Truncate(ctx context.Context) {
	if _, err := db.pool.Exec(ctx, "TRUNCATE TABLE reminders RESTART IDENTITY;"); err != nil {
		log.Fatal(err)
	}
}
