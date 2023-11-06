package model

import "time"

type HomeworkFromClass struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreateTime   time.Time `json:"createTime"`
	DeadlineTime time.Time `json:"deadlineTime"`
	File         string    `json:"file"`
}

type HomeworkListFromClass struct {
	Homeworks []*HomeworkFromClass `json:"homeworks,omitempty"`
}

type HomeworkByID struct {
	ClassID      int       `json:"classID"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreateTime   time.Time `json:"createTime"`
	DeadlineTime time.Time `json:"deadlineTime"`
	File         string    `json:"file"`
}

type HomeworkCreate struct {
	ClassID      int       `json:"classID"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	DeadlineTime time.Time `json:"deadlineTime"`
	File         string    `json:"file"`
}

type HomeworkCreateResponse struct {
	ID         int       `json:"id"`
	CreateTime time.Time `json:"createTime"`
}
