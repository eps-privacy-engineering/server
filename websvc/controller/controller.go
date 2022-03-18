package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/websvc/model"
	"server/websvc/response"
)

func GetWebsiteAttr(w http.ResponseWriter, r *http.Request) {
	//msg, _ := ioutil.ReadAll(r.Body)
	//req := new(request.GetWebsiteAttrRequest)
	//err := json.Unmarshal(msg, req)
	//if err != nil {
	//	return
	//}
	resp := new(response.SitePrivacyAttrResponse)
	bytes, err := ioutil.ReadFile("tmpDB.txt")
	if err != nil {
		return
	}
	ccpa := &model.CCPARights{}
	json.Unmarshal(bytes, ccpa)
	resp.CCPA = ccpa
	respJson, _ := json.Marshal(resp)
	fmt.Println("resp json", respJson)
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(respJson)
}
