package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"server/websvc/cache"
	"server/websvc/controller"
	"server/websvc/model"
)

func InitMemoryStore() {
	cache.LocalMap = make(map[string]*model.CCPARights, 0)
	bytes, err := ioutil.ReadFile("tmpDB.txt")
	if err != nil {
		return
	}
	json.Unmarshal(bytes, &cache.LocalMap)
}

func InitRouter() {
	http.HandleFunc("/get_website_attr", controller.GetWebsiteAttr)
	http.HandleFunc("/update_website_attr", controller.UpdateWebsiteAttr)
	http.HandleFunc("/extend_path_node", controller.ExtendPathNode)
}

func main() {
	InitMemoryStore()
	InitRouter()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
