package model

import "time"

type Homework struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreateTime   time.Time `json:"createTime"`
	DeadlineTime time.Time `json:"deadlineTime"`
	File         string    `json:"file"`
}

type HomeworkList struct {
	Homeworks []*Homework `json:"homeworks,omitempty"`
}

type HomeworkResponse struct {
	Homework Homework `json:"homework"`
}

type HomeworkByID struct {
	ClassID      int       `json:"classID"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreateTime   time.Time `json:"createTime"`
	DeadlineTime time.Time `json:"deadlineTime"`
	File         string    `json:"file"`
}

type HomeworkByIDResponse struct {
	Homework HomeworkByID `json:"homework"`
}

type HomeworkCreate struct {
	ClassID      int       `json:"classID"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	DeadlineTime time.Time `json:"deadlineTime"`
	File         string    `json:"file"`
}
