package model

type TeacherSignUp struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type TeacherProfile struct {
	Name string `json:"name"`
}

type TeacherProfileResponse struct {
	Teacher TeacherProfile `json:"teacher"`
}
