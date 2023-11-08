package model

import "time"

type Message struct {
	ID              int       `json:"id"`
	Text            string    `json:"text"`
	IsAuthorTeacher bool      `json:"ismine"`
	Attaches        []string  `json:"attaches,omitempty"`
	CreateTime      time.Time `json:"date"`
	IsRead          bool      `json:"isread"`
}
