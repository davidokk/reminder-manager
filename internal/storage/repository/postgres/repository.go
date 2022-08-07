package postgres

import "github.com/jackc/pgx/v4/pgxpool"

// Repository provides work with postgres
type Repository struct {
	pool *pgxpool.Pool
}

// NewRepository connects to db and initialized new Repository
func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool}
}

// Close closes db pool
func (r *Repository) Close() {
	r.pool.Close()
}
