package usecase

import (
	"log"
	"main/domain/model"
)

func (uc *Usecase) CreateSession(teacherLogin string) (*model.Session, error) {
	log.Println("call Create session", teacherLogin)
	return uc.store.CreateSession(teacherLogin)
}

func (uc *Usecase) CheckSession(in string) (string, error) {
	log.Println("call Check Session", in)
	return uc.store.CheckSession(in)
}

func (uc *Usecase) DeleteSession(in string) error {
	log.Println("call Delete Session", in)
	return uc.store.DeleteSession(in)
}
