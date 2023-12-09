package repository

import (
	e "main/domain/errors"
	"main/domain/model"
)

func (s *Store) AddTeacher(in *model.TeacherSignUp) error {
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

func (s *Store) GetTeacherProfile(id int) (*model.TeacherProfile, error) {
	var teacher model.TeacherProfile
	if err := s.db.QueryRow(
		`SELECT name FROM teachers WHERE id = $1;`,
		id,
	).Scan(&teacher.Name); err != nil {
		return nil, e.StacktraceError(err)
	}

	return &teacher, nil
}

func (s *Store) GetTeacherProfileByLoginDB(login string) (*model.TeacherDB, error) {
	var teacher model.TeacherDB

	if err := s.db.QueryRow(
		`SELECT id, password, name FROM teachers WHERE login = $1;`,
		login,
	).Scan(&teacher.ID, &teacher.Password, &teacher.Name); err != nil {
		return nil, e.StacktraceError(err)
	}

	teacher.Login = login
	return &teacher, nil
}
