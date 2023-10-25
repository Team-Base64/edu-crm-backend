package model

import "time"

type Error struct {
	Error interface{} `json:"error,omitempty"`
}

type Response struct {
	Body interface{} `json:"body,omitempty"`
}

type TeacherDB struct {
	ID        int     `json:"id"`
	Login     string  `json:"login"`
	Name      string  `json:"name"`
	Password  string  `json:"password"`
	TgAccount *string `json:"tgAccount,omitempty"`
	VkAccount *string `json:"vkAccount,omitempty"`
	TgBotLink *string `json:"tgBotLink,omitempty"`
	VkBotLink *string `json:"vkBotLink,omitempty"`
}

type StudentDB struct {
	InviteHash string  `json:"inviteHash"`
	Name       string  `json:"name"`
	ParentName *string `json:"parentName"`
	TgStudent  *string `json:"tgStudent,omitempty"`
	VkStudent  *string `json:"vkStudent,omitempty"`
	TgParent   *string `json:"tgParent,omitempty"`
	VkParent   *string `json:"vkParent,omitempty"`
}

type CreateStudentDB struct {
	Name string `json:"name"`
}

type ChatDB struct {
	ID          int    `json:"id"`
	TeacherID   int    `json:"teacherID"`
	StudentHash string `json:"studentHash"`
}

type CreateMessage struct {
	ChatID          int    `json:"chatid,omitempty"`
	Text            string `json:"text"`
	IsAuthorTeacher bool   `json:"isAuthorTeacher"`
}

type MessageChat struct {
	Text            string     `json:"text"`
	IsAuthorTeacher bool       `json:"isAuthorTeacher"`
	Attaches        *[]string  `json:"attaches"`
	Time            *time.Time `json:"time"`
}

type Chat struct {
	Messages []*MessageChat `json:"messages,omitempty"`
}

type ChatInfo struct {
	ChatID int `json:"chatid,omitempty"`
}

//	type Chats struct {
//		Chats []*Chat `json:"chats,omitempty"`
//	}
type Chats struct {
	Chats []*ChatInfo `json:"chats,omitempty"`
}

type MessageDB struct {
	ID              int        `json:"id,omitempty"`
	ChatID          int        `json:"chatid,omitempty"`
	Text            string     `json:"text"`
	IsAuthorTeacher bool       `json:"isAuthorTeacher"`
	Attaches        *[]string  `json:"attaches,omitempty"`
	Time            *time.Time `json:"time"`
}
