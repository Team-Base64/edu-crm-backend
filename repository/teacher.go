package repository

import (
	"database/sql"
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

func (s *Store) GetTeacherProfileByLoginDB(login string) (*model.TeacherDB, error) {
	var teacher model.TeacherDB
	rows, err := s.db.Query(`SELECT id, password, name FROM teachers WHERE login = $1;`, login)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()
	if err == sql.ErrNoRows {
		return nil, e.StacktraceError(e.ErrUnauthorized401)
	}
	for rows.Next() {
		err := rows.Scan(&teacher.ID, &teacher.Password, &teacher.Name)
		if err != nil {
			return nil, e.StacktraceError(err)
		}
	}
	teacher.Login = login
	return &teacher, nil
}
