package usecase

import (
	e "main/domain/errors"
	"main/domain/model"
)

func (uc *Usecase) CreateTeacher(params *model.TeacherSignUp) error {
	return uc.store.AddTeacher(params)
}

func (uc *Usecase) GetTeacherProfile(id int) (*model.TeacherProfile, error) {
	chat, err := uc.store.GetTeacherProfile(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return chat, nil
}
