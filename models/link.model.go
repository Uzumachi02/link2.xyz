package models

import "time"

// LinkModel ...
type LinkModel struct {
	ID        int
	TypeID    int
	URL       string
	TargetURL string
	Viewed    int
	CreatedAd time.Time
	UpdatedAt time.Time
}
