package repository

import (
	e "main/domain/errors"
	"main/domain/model"
)

func (s *Store) GetStudentByID(id int) (*model.StudentByID, error) {
	var stud model.StudentByID
	if err := s.db.QueryRow(
		`SELECT s.name, s.socialType, s.avatar, c.id FROM students s
		 JOIN chats c ON s.id = c.studentID
		 WHERE s.id = $1;`,
		id,
	).Scan(
		&stud.Name, &stud.SocialType, &stud.Avatar, &stud.ChatID,
	); err != nil {
		return nil, e.StacktraceError(err)
	}

	return &stud, nil
}

func (s *Store) GetStudentsFromClass(classID int) ([]model.StudentFromClass, error) {
	rows, err := s.db.Query(
		`SELECT s.id, s.name, s.socialType, s.avatar, c.id FROM students s
		 JOIN chats c ON s.id = c.studentID
		 WHERE c.classID = $1;`,
		classID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	students := []model.StudentFromClass{}
	for rows.Next() {
		var tmpStudent model.StudentFromClass

		if err := rows.Scan(
			&tmpStudent.ID,
			&tmpStudent.Name,
			&tmpStudent.SocialType,
			&tmpStudent.Avatar,
			&tmpStudent.ChatID,
		); err != nil {
			return nil, e.StacktraceError(err)
		}
		students = append(students, tmpStudent)
	}

	return students, nil
}
