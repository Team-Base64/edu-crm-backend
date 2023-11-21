package usecase

import (
	"log"
	e "main/domain/errors"
	"main/domain/model"

	"github.com/google/uuid"
)

func (uc *Usecase) CreateSession(teacherLogin string) (*model.Session, error) {
	log.Println("call Create session", teacherLogin)
	newUUID := uuid.New()
	//id := newUUID.String()
	sess := &model.Session{
		ID: newUUID.String(),
	}
	uc.mu.Lock()
	uc.sessions[sess.ID] = teacherLogin
	uc.mu.Unlock()

	return sess, nil
}

func (uc *Usecase) CheckSession(in string) (string, error) {
	log.Println("call Check Session", in)
	uc.mu.RLock()
	defer uc.mu.RUnlock()
	if sess, ok := uc.sessions[in]; ok {
		return sess, nil
	}
	return "", e.ErrUnauthorized401
}

func (uc *Usecase) DeleteSession(in string) error {
	log.Println("call Delete Session", in)
	uc.mu.Lock()
	defer uc.mu.Unlock()
	delete(uc.sessions, in)
	return nil
}
