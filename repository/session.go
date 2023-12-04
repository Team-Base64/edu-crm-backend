package repository

import (
	e "main/domain/errors"
	"main/domain/model"

	"github.com/google/uuid"
)

func (s *Store) CreateSession(teacherLogin string) (*model.Session, error) {
	newUUID := uuid.New()
	sess := &model.Session{
		ID: newUUID.String(),
	}
	if _, err := s.db.Exec(`INSERT INTO sessions (id, teacherLogin) VALUES ($1, $2);`, newUUID.String(), teacherLogin); err != nil {
		return nil, e.StacktraceError(err)
	}

	return sess, nil
}

func (s *Store) CheckSession(in string) (string, error) {
	teacherLogin := ""
	if err := s.db.QueryRow(`SELECT teacherLogin FROM sessions WHERE id = $1;`, in).Scan(&teacherLogin); err != nil {
		return "", e.StacktraceError(err)
	}
	return teacherLogin, nil
}

func (s *Store) DeleteSession(in string) error {
	if _, err := s.db.Exec(`DELETE FROM sessions WHERE id = $1;`, in); err != nil {
		return e.StacktraceError(err)
	}
	return nil
}
