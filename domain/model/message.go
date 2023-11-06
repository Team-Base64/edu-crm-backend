package model

import "time"

type Message struct {
	ID              int       `json:"id"`
	Text            string    `json:"text"`
	IsAuthorTeacher bool      `json:"ismine"`
	Attaches        []string  `json:"attaches,omitempty"`
	Time            time.Time `json:"time"`
	IsRead          bool      `json:"isread"`
}
