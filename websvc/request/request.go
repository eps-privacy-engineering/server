package request

import "server/websvc/model"

type FetchRequest struct {
	Mode int `json:"mode"`
}

type GetWebsiteAttrRequest struct {
	Host string `json:"host"`
}

type UpdateWebsiteAttrRequest struct {
	Host string `json:"host"`
	// ExerciseDetail: Refer to the model part
	ExerciseDetail *model.RightExercise `json:"exercise_detail"`
}
