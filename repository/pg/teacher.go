package pg

import (
	e "main/domain/errors"
	m "main/domain/model"
)

func (s *PostgreSqlStore) AddTeacher(in *m.TeacherSignUp) error {
	maxCount := 0
	if err := s.db.QueryRow(
		`SELECT count(*) FROM calendars;`).Scan(&maxCount); err != nil {
		return e.StacktraceError(err)
	}
	id := 0
	err := s.db.QueryRow(
		`INSERT INTO teachers (login, name, password) VALUES ($1, $2, $3) Returning id;`,
		in.Login, in.Name, in.Password,
	).Scan(&id)
	if err != nil {
		return e.StacktraceError(err)
	}
	if id > maxCount {
		if _, err := s.db.Exec(`DELETE FROM teachers WHERE id = $1;`, id); err != nil {
			return e.StacktraceError(err)
		}
		return e.StacktraceError(e.ErrServerError503)
	}

	return nil
}

func (s *PostgreSqlStore) GetTeacherProfile(id int) (*m.TeacherProfile, error) {
	var teacher m.TeacherProfile
	if err := s.db.QueryRow(
		`SELECT name FROM teachers WHERE id = $1;`,
		id,
	).Scan(&teacher.Name); err != nil {
		return nil, e.StacktraceError(err)
	}

	return &teacher, nil
}

func (s *PostgreSqlStore) GetTeacherProfileByLoginDB(login string) (*m.TeacherDB, error) {
	var teacher m.TeacherDB

	if err := s.db.QueryRow(
		`SELECT id, password, name FROM teachers WHERE login = $1;`,
		login,
	).Scan(&teacher.ID, &teacher.Password, &teacher.Name); err != nil {
		return nil, e.StacktraceError(err)
	}

	teacher.Login = login
	return &teacher, nil
}
