package store

import "github.com/Uzumachi02/link2.xyz/components/link"

// IStore ...
type IStore interface {
	Link() link.ILinkRepository
}
