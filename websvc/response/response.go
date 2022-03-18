package response

import "server/websvc/model"

type SitePrivacyAttrResponse struct {
	CCPA *model.CCPARights `json:"ccpa"`
}
