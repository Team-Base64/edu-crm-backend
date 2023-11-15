package model

import "time"

type Homework struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreateTime   time.Time `json:"createTime"`
	DeadlineTime time.Time `json:"deadlineTime"`
	Tasks        []int     `json:"tasks"`
}

type HomeworkList struct {
	Homeworks []Homework `json:"homeworks"`
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
	Tasks        []int     `json:"tasks"`
}

type HomeworkByIDResponse struct {
	Homework HomeworkByID `json:"homework"`
}

type HomeworkCreate struct {
	ClassID      int       `json:"classID"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	DeadlineTime time.Time `json:"deadlineTime"`
	Tasks        []int     `json:"tasks"`
}
