package pg

import (
	e "main/domain/errors"
	m "main/domain/model"
)

func (s *PostgreSqlStore) GetStudentByID(id int) (*m.StudentByID, error) {
	var stud m.StudentByID
	if err := s.db.QueryRow(
		`SELECT name, socialType, avatar FROM students WHERE id = $1;`,
		id,
	).Scan(
		&stud.Name, &stud.SocialType, &stud.Avatar,
	); err != nil {
		return nil, e.StacktraceError(err)
	}

	return &stud, nil
}

func (s *PostgreSqlStore) GetStudentsFromClass(classID int) ([]m.StudentFromClass, error) {
	rows, err := s.db.Query(
		`SELECT s.id, s.name, s.socialType, s.avatar, c.id FROM students s
		 JOIN chats c ON s.id = c.studentID
		 WHERE c.classID = $1;`,
		classID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	students := []m.StudentFromClass{}
	for rows.Next() {
		var tmpStudent m.StudentFromClass

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
