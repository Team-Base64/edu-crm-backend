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

type TeacherLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Session struct {
	ID string `json:"id"`
	//UserUUID string `json:"useruuid"`
}

type TeacherDB struct {
	ID       int    `json:"-"`
	Login    string `json:"login"`
	Name     string `json:"name"`
	Password string `json:"-"`
}
