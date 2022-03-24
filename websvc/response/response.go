package response

import "server/websvc/model"

type SitePrivacyAttrResponse struct {
	CCPA *model.CCPARights `json:"ccpa"`
}

type UpdateWebsiteAttrResponse struct {
	// Status: update result. 0 = ok, 1 = exist, 2= failure
	Status int `json:"status"`
}
