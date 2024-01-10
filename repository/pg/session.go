package pg

import (
	e "main/domain/errors"
	m "main/domain/model"

	"github.com/google/uuid"
)

func (s *PostgreSqlStore) CreateSession(teacherLogin string) (*m.Session, error) {
	newUUID := uuid.New()
	sess := &m.Session{
		ID: newUUID.String(),
	}
	if _, err := s.db.Exec(`INSERT INTO sessions (id, teacherLogin) VALUES ($1, $2);`, newUUID.String(), teacherLogin); err != nil {
		return nil, e.StacktraceError(err)
	}

	return sess, nil
}

func (s *PostgreSqlStore) CheckSession(in string) (string, error) {
	teacherLogin := ""
	if err := s.db.QueryRow(`SELECT teacherLogin FROM sessions WHERE id = $1;`, in).Scan(&teacherLogin); err != nil {
		return "", e.StacktraceError(err)
	}
	return teacherLogin, nil
}

func (s *PostgreSqlStore) DeleteSession(in string) error {
	if _, err := s.db.Exec(`DELETE FROM sessions WHERE id = $1;`, in); err != nil {
		return e.StacktraceError(err)
	}
	return nil
}
