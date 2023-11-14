package model

import "time"

type SolutionFromClass struct {
	ID         int       `json:"id"`
	HomeworkID int       `json:"hwID"`
	StudentID  int       `json:"studentID"`
	Text       string    `json:"text,omitempty"`
	CreateTime time.Time `json:"createTime"`
	File       string    `json:"file"`
}

type SolutionListFromClass struct {
	Solutions []*SolutionFromClass `json:"solutions"`
}

type SolutionForHw struct {
	ID         int       `json:"id"`
	StudentID  int       `json:"studentID"`
	Text       string    `json:"text,omitempty"`
	CreateTime time.Time `json:"createTime"`
	File       string    `json:"file"`
}

type SolutionListForHw struct {
	Solutions []*SolutionForHw `json:"solutions"`
}

type SolutionByID struct {
	HomeworkID int       `json:"hwID"`
	StudentID  int       `json:"studentID"`
	Text       string    `json:"text,omitempty"`
	CreateTime time.Time `json:"createTime"`
	File       string    `json:"file"`
}

type SolutionByIDResponse struct {
	Solution SolutionByID `json:"solution"`
}
