package model

import "time"

type ChatPreview struct {
	ChatID          int       `json:"chatid"`
	Name            string    `json:"name"`
	Img             string    `json:"cover"`
	IsRead          bool      `json:"isread"`
	LastMessageText string    `json:"lastmessagetext"`
	LastMessageDate time.Time `json:"lastmessagedate"`
}

type ChatPreviewList struct {
	Chats []*ChatPreview `json:"chats,omitempty"`
}

type Chat struct {
	Messages []*Message `json:"messages,omitempty"`
}
