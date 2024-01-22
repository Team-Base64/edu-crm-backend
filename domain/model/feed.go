package model

import "time"

type Post struct {
	ID         int       `json:"id"`
	Text       string    `json:"text"`
	Attaches   []string  `json:"attaches"`
	CreateTime time.Time `json:"createTime"`
}

type Feed struct {
	Posts []Post `json:"posts"`
}

type PostCreate struct {
	Text     string   `json:"text"`
	Attaches []string `json:"attaches"`
}

type PostResponse struct {
	Post Post `json:"post"`
}
