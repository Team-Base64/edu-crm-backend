package backend

import (
	"bytes"
	e "main/domain/errors"
	m "main/domain/model"

	"golang.org/x/crypto/argon2"
)

func (uc *BackendUsecase) HashPass(plainPassword string) []byte {
	salt := []byte(uc.saltString)
	hashedPass := argon2.IDKey([]byte(plainPassword), []byte(salt), 1, 64*1024, 4, 32)
	return append(salt, hashedPass...)
}

func (uc *BackendUsecase) CheckPass(passHash []byte, plainPassword string) bool {
	userPassHash := uc.HashPass(plainPassword)
	return bytes.Equal(userPassHash, passHash)
}

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
