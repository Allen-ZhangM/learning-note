package entity

type LastClickMd5Info struct {
	AccountID   int32  `json:"account_id"`
	CampaignID  string `json:"campaign_id"`
	AdgroupID   string `json:"adgroup_id"`
	AdID        string `json:"ad_id"`
	Dtu         string `json:"dtu"`
	Platform    string `json:"platform"`
	ChannelID   int32  `json:"channel_id"`
	CategoryID  string `json:"category_id"`
	ProductID   string `json:"product_id"`
	ExtensionID string `json:"extension_id"`
	ClickTime   string `json:"click_time"`
}
