package postgres

import (
	"log"

	"github.com/pashagolub/pgxmock"
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

func (f *repositoryFixtures) tearDown() {
	f.repository.Close()
}
