package link

type createLinkModel struct {
	TargetURL string `json:"target_url" form:"target_url" query:"target_url"`
	Content   string `json:"content" form:"content" query:"content"`
	Type      int    `json:"type" form:"type" query:"type"`
}
