package postgres

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type poolInterface interface {
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)

	Close()
}

// Repository provides work with postgres
type Repository struct {
	pool poolInterface
}

// NewRepository initialized new Repository with pool
func NewRepository(pool poolInterface) *Repository {
	return &Repository{pool}
}

// Close closes db pool
func (r *Repository) Close() {
	r.pool.Close()
}
