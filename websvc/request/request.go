package request

import "server/websvc/model"

type FetchRequest struct {
	Mode int `json:"mode"`
}

type GetWebsiteAttrRequest struct {
	Host string `json:"host"`
}

type FinishPathRequest struct {
	Host      string `json:"host"`
	RightType string `json:"right_type"`
}

type UpdateWebsiteAttrRequest struct {
	Host string `json:"host"`
	// ExerciseDetail: Refer to the model part
	ExerciseDetail *model.RightExercise `json:"exercise_detail"`
}

type ExtendPathNodeRequest struct {
	Host      string              `json:"host"`
	Node      *model.ExerciseNode `json:"node"`
	RightType string              `json:"right_type"`
}
