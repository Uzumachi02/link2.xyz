package models

import "time"

// TypeModel ...
type TypeModel struct {
	ID        int
	Name      string
	Status    int
	CreatedAd time.Time
}
