package model

import "time"

type SolutionFromClass struct {
	ID        int       `json:"id"`
	HwID      int       `json:"hwID"`
	StudentID int       `json:"studentID"`
	Text      string    `json:"text,omitempty"`
	Time      time.Time `json:"time"`
	File      string    `json:"file"`
}

type SolutionsFromClass struct {
	Solutions []*SolutionFromClass `json:"solutions"`
}

type SolutionForHw struct {
	ID        int       `json:"id"`
	StudentID int       `json:"studentID"`
	Text      string    `json:"text,omitempty"`
	Time      time.Time `json:"time"`
	File      string    `json:"file"`
}

type SolutionsForHw struct {
	Solutions []*SolutionForHw `json:"solutions"`
}

type SolutionByID struct {
	HwID      int       `json:"hwID"`
	StudentID int       `json:"studentID"`
	Text      string    `json:"text,omitempty"`
	Time      time.Time `json:"time"`
	File      string    `json:"file"`
}
