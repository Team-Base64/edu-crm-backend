package model

type StudentFromClass struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	SocialType string `json:"socialType"`
	ChatID     int    `json:"chatID"`
	Avatar     string `json:"avatar"`
}

type StudentListFromClass struct {
	Students []StudentFromClass `json:"students"`
}

type StudentByID struct {
	Name       string `json:"name"`
	SocialType string `json:"socialType"`
	Avatar     string `json:"avatar"`
}

type StudentByIDResponse struct {
	Student StudentByID `json:"student"`
}
