package link

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

type createLinkModel struct {
	TargetURL string `json:"target_url" form:"target_url" query:"target_url"`
	Content   string `json:"content" form:"content" query:"content"`
}
