package model

type StudentFromClass struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	SocialType string `json:"socialType"`
}

type StudentListFromClass struct {
	Students []*StudentFromClass `json:"students"`
}

type StudentByID struct {
	Name       string `json:"name"`
	SocialType string `json:"socialType"`
}
