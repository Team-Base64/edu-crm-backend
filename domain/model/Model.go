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
	IsAuthorTeacher bool   `json:"ismine"`
}

type MessageChat struct {
	ID              int        `json:"id"`
	Text            string     `json:"text"`
	IsAuthorTeacher bool       `json:"ismine"`
	ChatID          int        `json:"chatid"`
	Attaches        *[]string  `json:"attaches,omitempty"`
	Time            *time.Time `json:"time"`
	IsRead          bool       `json:"isread"`
}

type Chat struct {
	Messages []*MessageChat `json:"messages,omitempty"`
}

type ChatInfo struct {
	ChatID          int        `json:"chatid"`
	Name            string     `json:"name"`
	Img             string     `json:"cover"`
	IsRead          bool       `json:"isread"`
	LastMessageText string     `json:"lastmessagetext"`
	LastMessageDate *time.Time `json:"lastmessagedate"`
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
	IsRead          bool       `json:"isread"`
}
