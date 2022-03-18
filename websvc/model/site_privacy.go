package model

type SitePrivacyAttr struct {
	Host string     `json:"host"`
	CCPA CCPARights `json:"ccpa_rights"`
}
