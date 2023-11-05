package model

type ClassInfo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	InviteToken string `json:"inviteToken"`
}

type ClassesInfo struct {
	Classes []*ClassInfo `json:"classes,omitempty"`
}
