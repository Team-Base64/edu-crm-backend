package model

type Student struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	SocialType string `json:"socialType"`
}

type StudentsFromClass struct {
	Students []*Student
}
