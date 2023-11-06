package model

import "time"

type Post struct {
	ID       int       `json:"id"`
	Text     string    `json:"text"`
	Attaches []string  `json:"attaches,omitempty"`
	Time     time.Time `json:"time"`
}

type Feed struct {
	Posts []*Post `json:"posts"`
}

type PostCreate struct {
	Text     string   `json:"text"`
	Attaches []string `json:"attaches,omitempty"`
}

type PostCreateResponse struct {
	ID   int       `json:"id"`
	Time time.Time `json:"time"`
}
