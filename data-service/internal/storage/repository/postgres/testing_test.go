package postgres

import (
	"github.com/pashagolub/pgxmock"
	"log"
	"testing"
)

type repositoryFixtures struct {
	repository *Repository
	pool       pgxmock.PgxPoolIface
}

func setUp() repositoryFixtures {
	var f repositoryFixtures
	var err error
	if f.pool, err = pgxmock.NewPool(); err != nil {
		log.Fatal(err.Error())
	}
	f.repository = NewRepository(f.pool)

	return f
}

func (f *repositoryFixtures) tearDown(t *testing.T) {
	if err := f.pool.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %s", err)
	}
	f.repository.Close()
}
