package model

import "time"

type ChatPreview struct {
	ChatID          int       `json:"chatID"`
	Name            string    `json:"name"`
	Img             string    `json:"cover"`
	SocialType      string    `json:"socialType"`
	IsRead          bool      `json:"isread"`
	LastMessageText string    `json:"text"`
	LastMessageDate time.Time `json:"date"`
}

type ChatPreviewList struct {
	Chats []*ChatPreview `json:"chats,omitempty"`
}

type Chat struct {
	Messages []*Message `json:"messages,omitempty"`
}
