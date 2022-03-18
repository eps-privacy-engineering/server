package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/websvc/cache"
	"server/websvc/request"
	"server/websvc/response"
)

func GetWebsiteAttr(w http.ResponseWriter, r *http.Request) {
	msg, _ := ioutil.ReadAll(r.Body)
	req := new(request.GetWebsiteAttrRequest)
	err := json.Unmarshal(msg, req)
	if err != nil {
		return
	}
	fmt.Println(req.Host)
	resp := new(response.SitePrivacyAttrResponse)
	fmt.Println(cache.LocalMap)
	ccpa := cache.LocalMap[req.Host]
	resp.CCPA = ccpa
	respJson, _ := json.Marshal(resp)
	fmt.Println("resp json", respJson)
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(respJson)
}
