package models

import "time"

// AgentModel ...
type AgentModel struct {
	ID        int
	Hash      string
	UserAgent string
	Platform  string
	Browser   string
	Version   string
	CreatedAd time.Time
}
