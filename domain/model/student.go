package model

type Student struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	SocialType string `json:"socialType"`
}

type StudentListFromClass struct {
	Students []*Student `json:"students"`
}
