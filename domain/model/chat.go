package model

import "time"

type ChatPreview struct {
	ChatID          int       `json:"chatID"`
	StudentID       int       `json:"studentID"`
	StudentName     string    `json:"studentName"`
	StudentAvatar   string    `json:"studentAvatar"`
	SocialType      string    `json:"socialType"`
	IsRead          bool      `json:"isread"`
	LastMessageText string    `json:"text"`
	LastMessageDate time.Time `json:"date"`
}

type ChatPreviewList struct {
	Chats []ChatPreview `json:"chats"`
}

type Chat struct {
	Messages []Message `json:"messages"`
}
