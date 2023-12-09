package model

type Task struct {
	ID          int      `json:"id"`
	Description string   `json:"description"`
	Attaches    []string `json:"attaches"`
}

type TaskCreate struct {
	Description string   `json:"description"`
	Attaches    []string `json:"attaches"`
}

type TaskByID struct {
	Description string   `json:"description"`
	Attaches    []string `json:"attaches"`
}

type TaskListByTeacherID struct {
	Tasks []Task `json:"tasks"`
}

type TaskCreateResponse struct {
	ID int `json:"id"`
}
type TaskByIDResponse struct {
	Task TaskByID `json:"task"`
}
