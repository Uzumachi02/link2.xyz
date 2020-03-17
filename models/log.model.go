package models

import "time"

// LogModel ...
type LogModel struct {
	ID        int
	LinkID    int
	IPID      int
	AgentID   int
	Others    string
	CreatedAd time.Time
}
