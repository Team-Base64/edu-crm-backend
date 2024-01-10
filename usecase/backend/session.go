package backend

import (
	"log"

	m "main/domain/model"
)

func (uc *BackendUsecase) CreateSession(teacherLogin string) (*m.Session, error) {
	log.Println("call Create session", teacherLogin)
	return uc.dataStore.CreateSession(teacherLogin)
}

func (uc *BackendUsecase) CheckSession(in string) (string, error) {
	log.Println("call Check Session", in)
	return uc.dataStore.CheckSession(in)
}

func (uc *BackendUsecase) DeleteSession(in string) error {
	log.Println("call Delete Session", in)
	return uc.dataStore.DeleteSession(in)
}
