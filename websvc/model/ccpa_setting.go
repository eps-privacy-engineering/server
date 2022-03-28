package model

type CCPARights struct {
	CCPADoNotSell *RightExercise `json:"ccpa_do_not_sell"`
	CCPADelete    *RightExercise `json:"ccpa_delete"`
	CCPACopy      *RightExercise `json:"ccpa_copy"`
}

func NewCCPARights(dns, d, c *RightExercise) *CCPARights {
	ccpa := &CCPARights{
		CCPADoNotSell: dns,
		CCPADelete:    d,
		CCPACopy:      c,
	}
	return ccpa
}

type RightExercise struct {
	RightType    string          `json:"right_type"`
	ExercisePath []*ExerciseNode `json:"exercise_path"`
}

func NewRightExercise(rightType string, pathList []*ExerciseNode) *RightExercise {
	path := make([]*ExerciseNode, 0)
	for i := 0; i < len(pathList); i++ {
		path = append(path, pathList[i])
	}
	re := &RightExercise{
		RightType:    rightType,
		ExercisePath: pathList,
	}
	return re
}

type ExerciseNode struct {
	// Page: url without params
	Page string `json:"page"`
	// Text: html element text
	Text string `json:"text"`
	// Category: html element category
	Category string `json:"category"`
	// Op: click
	OperationType string `json:"operation_type"`
	// HTMLID:
	HTMLID string `json:"html_id"`
}

func NewExerciseNode(text, category, operationType string) *ExerciseNode {
	en := &ExerciseNode{
		Text:          text,
		Category:      category,
		OperationType: operationType,
	}
	return en
}
