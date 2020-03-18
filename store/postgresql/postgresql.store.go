package postgresql

import (
	"log"

	"github.com/Uzumachi02/link2.xyz/components/link"
	"github.com/go-pg/pg/v9"
)

var (
	pgStore *PostgresqlStore
)

// PostgresqlStore ...
type PostgresqlStore struct {
	databaseURL    string
	db             *pg.DB
	linkRepository *LinkRepository
}

// Connect ...
func (ps *PostgresqlStore) Connect() error {
	options, err := pg.ParseURL(ps.databaseURL)
	if err != nil {
		return err
	}

	db := pg.Connect(options)

	_, err = db.Exec("SELECT 1")
	if err != nil {
		return err
	}

	ps.db = db
	return nil
}

// Close ...
func (ps *PostgresqlStore) Close() {
	if ps.db != nil {
		ps.db.Close()
	}
}

// Link ...
func (ps *PostgresqlStore) Link() link.ILinkRepository {
	if ps.linkRepository != nil {
		return ps.linkRepository
	}

	ps.linkRepository = &LinkRepository{
		store: ps,
	}

	return ps.linkRepository
}

// NewStore ...
func NewStore(databaseURL string) *PostgresqlStore {

	// new store
	s := &PostgresqlStore{
		databaseURL: databaseURL,
	}

	// configure the store
	if err := s.Connect(); err != nil {
		log.Fatal(err)
	}

	pgStore = s

	return s
}

// GetStore ...
func GetStore() *PostgresqlStore {
	if pgStore != nil {
		return pgStore
	}

	return nil
}
