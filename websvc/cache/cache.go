package cache

import "server/websvc/model"

var LocalMap map[string]*model.CCPARights

func LocalMapUpdate(key string, ccpaRights *model.CCPARights) {
	LocalMap[key] = ccpaRights
}
