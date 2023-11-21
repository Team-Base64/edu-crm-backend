package model

import "time"

type SolutionFromClass struct {
	ID                int       `json:"id"`
	HomeworkID        int       `json:"hwID"`
	StudentID         int       `json:"studentID"`
	Text              string    `json:"text"`
	CreateTime        time.Time `json:"createTime"`
	File              string    `json:"file"`
	Status            string    `json:"status"` // new | approve | reject
	TeacherEvaluation string    `json:"teacherEvaluation"`
}

type SolutionListFromClass struct {
	Solutions []SolutionFromClass `json:"solutions"`
}

type SolutionForHw struct {
	ID                int       `json:"id"`
	StudentID         int       `json:"studentID"`
	Text              string    `json:"text"`
	CreateTime        time.Time `json:"createTime"`
	File              string    `json:"file"`
	Status            string    `json:"status"` // new | approve | reject
	TeacherEvaluation string    `json:"teacherEvaluation"`
}

type SolutionListForHw struct {
	Solutions []SolutionForHw `json:"solutions"`
}

type SolutionByID struct {
	HomeworkID        int       `json:"hwID"`
	StudentID         int       `json:"studentID"`
	Text              string    `json:"text"`
	CreateTime        time.Time `json:"createTime"`
	File              string    `json:"file"`
	Status            string    `json:"status"` // new | approve | reject
	TeacherEvaluation string    `json:"teacherEvaluation"`
}

type SolutionByIDResponse struct {
	Solution SolutionByID `json:"solution"`
}

type SolutionInfoForEvaluationMsg struct {
	HomeworkTitle      string
	SolutionCreateTime time.Time
}
