package model

type ClassInfo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	InviteToken string `json:"inviteToken"`
}

type ClassInfoResponse struct {
	Class ClassInfo `json:"class"`
}

type ClassInfoList struct {
	Classes []*ClassInfo `json:"classes,omitempty"`
}

type ClassCreate struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}
