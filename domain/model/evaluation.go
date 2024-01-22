package model

type TaskEvaluation struct {
	TaskID     int    `json:"id"`
	Evaluation string `json:"evaluation"`
}

type SolutionEvaluation struct {
	IsApproved bool             `json:"isApproved"`
	Tasks      []TaskEvaluation `json:"tasks"`
}
