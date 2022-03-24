package cache

import (
	"encoding/json"
	"io/ioutil"
	"server/websvc/model"
)

var LocalMap map[string]*model.CCPARights

func LocalMapUpdate(key string, ccpaRights *model.CCPARights) {
	LocalMap[key] = ccpaRights
	jsBytes, _ := json.MarshalIndent(LocalMap, "", "\t")
	err := ioutil.WriteFile("tmpDB.txt", jsBytes, 0644)
	if err != nil {
		panic(err)
	}
}
