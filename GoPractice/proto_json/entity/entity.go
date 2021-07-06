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

type SuperlinkClickSecondCache struct {
	AccountID    int    `json:"account_id"`
	CampaignID   string `json:"campaign_id"`
	AdgroupID    string `json:"adgroup_id"`
	AdID         string `json:"ad_id"`
	Dtu          string `json:"dtu"` //--
	Platform     string `json:"platform"`
	ChannelID    int    `json:"channel_id"`
	CategoryID   string `json:"category_id"` //--
	ProductID    string `json:"product_id"`
	ExtensionID  string `json:"extension_id"`
	TuID         string `json:"tuid"`
	OaID         string `json:"oaid"`
	DeviceMd5    string `json:"device_md5"`
	ClickTime    string `json:"click_time"`
	AndroidID    string `json:"android_id"`
	AndroidIDMd5 string `json:"android_id_md5"`
}

type ClickFirstCache struct {
	Md5Key    string      `json:"md5_key"`
	ClickTime interface{} `json:"click_time"`
	Time      int64       `json:"time"`
	Platform  string      `json:"platform"`
}

type ClickFirstCacheUpdate struct {
	Md5Key    string      `json:"md5_key"`
	ClickTime interface{} `json:"click_time"`
	Time      int64       `json:"time"`
}

type ClickFirstCacheUpdate2 struct {
	Md5Key string `json:"md5_key"`
	Time   int64  `json:"time"`
}
