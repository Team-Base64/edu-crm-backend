package repository

import (
	e "main/domain/errors"
	"main/domain/model"
)

func (s *Store) CheckClassExistence(id int) error {
	var tmp int
	if err := s.db.QueryRow(
		`SELECT 1 FROM classes WHERE id = $1;`,
		id,
	).Scan(&tmp); err != nil {
		return e.StacktraceError(err)
	}
	return nil
}

func (s *Store) AddClass(teacherID int, inviteToken string, newClass *model.ClassCreate) (int, error) {
	var id int
	if err := s.db.QueryRow(
		`INSERT INTO classes (teacherID, title, description, inviteToken)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id;`,
		teacherID, newClass.Title, newClass.Description, inviteToken,
	).Scan(&id); err != nil {
		return 0, e.StacktraceError(err)
	}

	return int(id), nil
}

func (s *Store) GetClassByID(id int) (*model.ClassInfo, error) {
	var class model.ClassInfo
	if err := s.db.QueryRow(
		`SELECT title, description, inviteToken FROM classes WHERE id = $1;`,
		id,
	).Scan(
		&class.Title,
		&class.Description,
		&class.InviteToken,
	); err != nil {
		return nil, e.StacktraceError(err)
	}

	class.ID = id
	return &class, nil
}

func (s *Store) GetClassesByID(teacherID int) (*model.ClassInfoList, error) {
	rows, err := s.db.Query(
		`SELECT id, title, description, inviteToken FROM classes WHERE teacherID = $1;`,
		teacherID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	classes := []*model.ClassInfo{}
	for rows.Next() {
		var tmpClass model.ClassInfo

		if err := rows.Scan(
			&tmpClass.ID, &tmpClass.Title,
			&tmpClass.Description, &tmpClass.InviteToken,
		); err != nil {
			return nil, e.StacktraceError(err)
		}
		classes = append(classes, &tmpClass)
	}

	return &model.ClassInfoList{Classes: classes}, nil
}
