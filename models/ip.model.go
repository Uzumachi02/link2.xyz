package models

import "time"

// IPModel ...
type IPModel struct {
	ID          int
	IP          string
	CountryCode string
	CountryName string
	CreatedAd   time.Time
}
