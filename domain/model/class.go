package model

type ClassDB struct {
	ID          int
	TeacherID   int
	Title       string
	Description string
	InviteToken string
}

type Class struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	InviteToken string `json:"inviteToken,omitempty"`
}

type Classes struct {
	Classes []*Class `json:"classes,omitempty"`
}
