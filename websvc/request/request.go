package request

type FetchRequest struct {
	Mode int `json:"mode"`
}

type GetWebsiteAttrRequest struct {
	Host string `json:"host"`
}
