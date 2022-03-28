package response

import "server/websvc/model"

type SitePrivacyAttrResponse struct {
	Status int               `json:"status"` // 0 = ok, 1 = failure
	CCPA   *model.CCPARights `json:"ccpa"`
}

type UpdateWebsiteAttrResponse struct {
	// Status: update result. 0 = ok, 1 = exist, 2= failure
	Status int `json:"status"`
}
