package repository

import (
	e "main/domain/errors"
	"main/domain/model"
)

func (s *Store) GetStudentByID(id int) (*model.StudentByID, error) {
	var stud model.StudentByID
	if err := s.db.QueryRow(
		`SELECT name, socialType FROM students WHERE id = $1;`,
		id,
	).Scan(
		&stud.Name, &stud.SocialType,
	); err != nil {
		return nil, e.StacktraceError(err)
	}

	return &stud, nil
}

func (s *Store) GetStudentsFromClass(classID int) (*model.StudentListFromClass, error) {
	rows, err := s.db.Query(
		`SELECT s.id, s.name, s.socialType FROM students s
		 JOIN classes_students cs ON s.id = cs.studentID
		 WHERE cs.classID = $1;`,
		classID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	students := []*model.StudentFromClass{}
	for rows.Next() {
		var tmpStudent model.StudentFromClass

		if err := rows.Scan(
			&tmpStudent.ID,
			&tmpStudent.Name,
			&tmpStudent.SocialType,
		); err != nil {
			return nil, e.StacktraceError(err)
		}
		students = append(students, &tmpStudent)
	}

	return &model.StudentListFromClass{Students: students}, nil
}
