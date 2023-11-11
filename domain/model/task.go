package model

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description,omitempty"`
	Attach      string `json:"attach,omitempty"`
}

type TaskCreate struct {
	Description string `json:"description,omitempty"`
	Attach      string `json:"attach,omitempty"`
}

type TaskByID struct {
	Description string `json:"description,omitempty"`
	Attach      string `json:"attach,omitempty"`
}
