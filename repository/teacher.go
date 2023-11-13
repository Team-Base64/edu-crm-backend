package repository

import (
	e "main/domain/errors"
	"main/domain/model"
)

func (s *Store) AddTeacher(in *model.TeacherSignUp) error {
	_, err := s.db.Exec(
		`INSERT INTO teachers (login, name, password) VALUES ($1, $2, $3);`,
		in.Login, in.Name, in.Password,
	)
	if err != nil {
		return e.StacktraceError(err)
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
