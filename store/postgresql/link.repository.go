package postgresql

import (
	"log"

	"github.com/Uzumachi02/link2.xyz/model"
	"github.com/Uzumachi02/link2.xyz/store"
	"github.com/go-pg/pg/v9"
)

// LinkRepository ...
type LinkRepository struct {
	store *PostgresqlStore
}

// Create ...
func (l *LinkRepository) Create(*model.Link) error {
	return nil
}

// Find ...
func (l *LinkRepository) Find(int) (*model.Link, error) {
	return nil, nil
}

// HasUniqid ...
func (l *LinkRepository) HasUniqid(uniqid string) (bool, error) {
	var has bool
	_, err := l.store.db.QueryOne(pg.Scan(&has), "SELECT EXISTS(SELECT 1 FROM links WHERE url = ? LIMIT 1)", uniqid)
	if err != nil {
		return false, err
	}

	log.Println(has)

	return has, nil
}

// NewLinkRepository ...
func NewLinkRepository(store ...*PostgresqlStore) store.ILinkRepository {
	r := new(LinkRepository)

	if len(store) > 0 {
		r.store = store[0]
	} else {
		r.store = GetStore()
	}

	return r
}
