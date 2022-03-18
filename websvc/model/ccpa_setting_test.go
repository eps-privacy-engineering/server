package model

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestNewCCPARights(t *testing.T) {
	eNode := NewExerciseNode("Do Not Sell My Personal Information", "href", "click")
	ep := make([]*ExerciseNode, 0)
	ep = append(ep, eNode)
	re := NewRightExercise("CCPA_Sell", ep)
	ccpa := NewCCPARights(re, nil, nil)
	localMap := make(map[string]*CCPARights, 0)
	localMap["www.xfinity.com"] = ccpa
	jsBytes, _ := json.MarshalIndent(localMap, "", "\t")
	err := ioutil.WriteFile("tmpDB.txt", jsBytes, 0644)
	if err != nil {
		panic(err)
	}
}
