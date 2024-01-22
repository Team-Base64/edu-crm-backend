package model

type ClassBroadcastMessage struct {
	ClassID     int
	Title       string
	Description string
	Attaches    []string
}

type SingleMessage struct {
	ChatID   int
	Text     string
	Attaches []string
}
