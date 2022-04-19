package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/websvc/cache"
	"server/websvc/model"
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
	resp := new(response.SitePrivacyAttrResponse)
	ccpa, ok := cache.LocalMap[req.Host]
	if !ok {
		resp.Status = 1
		resp.CCPA = nil
	} else {
		resp.CCPA = ccpa
	}
	respJson, _ := json.Marshal(resp)
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(respJson)
}

func UpdateWebsiteAttr(w http.ResponseWriter, r *http.Request) {
	msg, _ := ioutil.ReadAll(r.Body)
	req := new(request.UpdateWebsiteAttrRequest)
	err := json.Unmarshal(msg, req)
	if err != nil {
		return
	}
	resp := new(response.UpdateWebsiteAttrResponse)
	ccpa, ok := cache.LocalMap[req.Host]
	if ok {
		resp.Status = 1
	} else {
		ccpa = new(model.CCPARights)
		if req.ExerciseDetail.RightType == "CCPADelete" {
			ccpa.CCPADelete = req.ExerciseDetail
			resp.Status = 0
		} else if req.ExerciseDetail.RightType == "CCPACopy" {
			ccpa.CCPACopy = req.ExerciseDetail
			resp.Status = 0
		} else if req.ExerciseDetail.RightType == "CCPADoNotSell" {
			ccpa.CCPADoNotSell = req.ExerciseDetail
			resp.Status = 0
		} else {
			resp.Status = 2
		}
	}
	cache.LocalMapUpdate(req.Host, ccpa)
	respJson, _ := json.Marshal(resp)
	fmt.Println("resp json", respJson)
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(respJson)
}

func ExtendPathNode(w http.ResponseWriter, r *http.Request) {
	msg, _ := ioutil.ReadAll(r.Body)
	req := new(request.ExtendPathNodeRequest)
	err := json.Unmarshal(msg, req)
	if err != nil {
		return
	}
	resp := new(response.ExtendPathNodeResponse)
	ccpa, ok := cache.LocalMap[req.Host]
	if ok {
		resp.Status = 1
	} else {
		ccpa = new(model.CCPARights)
		if req.RightType == "CCPADelete" {
			ccpa.CCPADelete.ExercisePath = append(ccpa.CCPADelete.ExercisePath, req.Node)
			resp.Status = 0
		} else if req.RightType == "CCPACopy" {
			ccpa.CCPACopy.ExercisePath = append(ccpa.CCPACopy.ExercisePath, req.Node)
			resp.Status = 0
		} else if req.RightType == "CCPADoNotSell" {
			ccpa.CCPADoNotSell.ExercisePath = append(ccpa.CCPADoNotSell.ExercisePath, req.Node)
			resp.Status = 0
		} else {
			resp.Status = 2
		}
	}
	cache.LocalMapUpdate(req.Host, ccpa)
	respJson, _ := json.Marshal(resp)
	fmt.Println("resp json", respJson)
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(respJson)
}

func FinishPath(w http.ResponseWriter, r *http.Request) {
	msg, _ := ioutil.ReadAll(r.Body)
	req := new(request.FinishPathRequest)
	err := json.Unmarshal(msg, req)
	if err != nil {
		return
	}
	resp := new(response.UpdateWebsiteAttrResponse)
	ccpa, ok := cache.LocalMap[req.Host]
	if !ok {
		resp.Status = 1
	} else {
		if req.RightType == "CCPADelete" {
			if ccpa.CCPADelete == nil {
				resp.Status = 1
			} else {
				ccpa.CCPADelete.Finish = true
				resp.Status = 0
			}
		} else if req.RightType == "CCPACopy" {
			if ccpa.CCPACopy == nil {
				resp.Status = 1
			} else {
				ccpa.CCPACopy.Finish = true
				resp.Status = 0
			}
		} else if req.RightType == "CCPADoNotSell" {
			if ccpa.CCPADoNotSell == nil {
				resp.Status = 1
			} else {
				ccpa.CCPADoNotSell.Finish = true
				resp.Status = 0
			}
		}
	}
	cache.LocalMapUpdate(req.Host, ccpa)
	respJson, _ := json.Marshal(resp)
	fmt.Println("resp json", respJson)
	w.Header().Set("content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(respJson)
}
