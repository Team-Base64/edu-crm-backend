package backend

import (
	e "main/domain/errors"
	m "main/domain/model"
)

func (uc *BackendUsecase) CreateTeacher(params *m.TeacherSignUp) error {
	return uc.dataStore.AddTeacher(params)
}

func (uc *BackendUsecase) GetTeacherProfile(id int) (*m.TeacherProfile, error) {
	chat, err := uc.dataStore.GetTeacherProfile(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return chat, nil
}

func (uc *BackendUsecase) GetTeacherProfileByLogin(login string) (*m.TeacherDB, error) {
	chat, err := uc.dataStore.GetTeacherProfileByLoginDB(login)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return chat, nil
}

func (uc *BackendUsecase) SignUpTeacher(req *m.TeacherSignUp) error {
	err := uc.dataStore.AddTeacher(req)
	if err != nil {
		return e.StacktraceError(err)
	}
	return nil
}
