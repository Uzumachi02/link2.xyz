package postgresql

import "github.com/Uzumachi02/link2.xyz/components/link"

// LinkRepository ...
type LinkRepository struct {
	store *PostgresqlStore
}

// Create ...
func (l *LinkRepository) Create(*link.LinkModel) error {
	return nil
}

// Find ...
func (l *LinkRepository) Find(int) (*link.LinkModel, error) {
	return nil, nil
}

// NewLinkRepository ...
func NewLinkRepository(store ...*PostgresqlStore) link.ILinkRepository {
	r := new(LinkRepository)

	if len(store) > 0 {
		r.store = store[0]
	} else {
		r.store = GetStore()
	}

	return r
}
