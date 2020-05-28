package store

import "github.com/Uzumachi02/link2.xyz/model"

// ILinkRepository ...
type ILinkRepository interface {
	Create(*model.Link) error
	Find(int) (*model.Link, error)
	HasUniqid(string) (bool, error)
}
