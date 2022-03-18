package main

import (
	"net/http"
	"server/websvc/controller"
)

func InitRouter() {
	http.HandleFunc("/get_website_attr", controller.GetWebsiteAttr)
}

func main() {
	InitRouter()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
