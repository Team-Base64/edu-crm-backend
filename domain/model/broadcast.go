package model

import "time"

type ClassBroudcastMessage struct {
	ClassID      int       `json:"classID"`
	Title        string    `json:"title"`
	Description  string    `json:"description,omitempty"`
	DeadlineTime time.Time `json:"deadlineTime,omitempty"`
	Attaches     []string  `json:"attaches,omitempty"`
}
