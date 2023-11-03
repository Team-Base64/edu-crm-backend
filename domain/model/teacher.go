package model

type TeacherSignUp struct {
	Login string `json:"login"`
	Name  string `json:"name"`
}

type TeacherProfile struct {
	Name string `json:"name"`
}
