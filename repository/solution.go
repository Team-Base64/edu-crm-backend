package repository

import (
	"database/sql"
	e "main/domain/errors"
	"main/domain/model"
)

func makeStatus(statusFromDB sql.NullBool) string {
	if statusFromDB.Valid {
		if statusFromDB.Bool {
			return "approve"
		} else {
			return "reject"
		}
	}
	return "new"
}

func (s *Store) GetSolutionByID(id int) (*model.SolutionByID, error) {
	var sol model.SolutionByID
	var isApproved sql.NullBool

	if err := s.db.QueryRow(
		`SELECT homeworkID, studentID, text, createTime, file, isApproved, teacherEvaluation FROM solutions WHERE id = $1;`,
		id,
	).Scan(
		&sol.HomeworkID, &sol.StudentID,
		&sol.Text, &sol.CreateTime, &sol.File, &isApproved, &sol.TeacherEvaluation,
	); err != nil {
		return nil, e.StacktraceError(err)
	}
	sol.Status = makeStatus(isApproved)

	return &sol, nil
}

func (s *Store) GetSolutionsByClassID(classID int) (*model.SolutionListFromClass, error) {
	rows, err := s.db.Query(
		`SELECT s.id, s.homeworkID, s.studentID, s.text, s.createTime, s.file, s.isApproved, s.teacherEvaluation
		 FROM solutions s
		 JOIN homeworks h ON s.homeworkID = h.id
		 WHERE h.classID = $1;`,
		classID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	sols := []*model.SolutionFromClass{}
	for rows.Next() {
		var tmpSol model.SolutionFromClass
		var isApproved sql.NullBool

		if err := rows.Scan(
			&tmpSol.ID, &tmpSol.HomeworkID, &tmpSol.StudentID,
			&tmpSol.Text, &tmpSol.CreateTime, &tmpSol.File,
			&isApproved, &tmpSol.TeacherEvaluation,
		); err != nil {
			return nil, e.StacktraceError(err)
		}
		tmpSol.Status = makeStatus(isApproved)

		sols = append(sols, &tmpSol)
	}

	return &model.SolutionListFromClass{Solutions: sols}, nil
}

func (s *Store) GetSolutionsByHomeworkID(homeworkID int) (*model.SolutionListForHw, error) {
	rows, err := s.db.Query(
		`SELECT id, studentID, text, createTime, file, isApproved, teacherEvaluation FROM solutions WHERE homeworkID = $1;`,
		homeworkID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	sols := []*model.SolutionForHw{}
	for rows.Next() {
		var tmpSol model.SolutionForHw
		var isApproved sql.NullBool

		if err := rows.Scan(
			&tmpSol.ID, &tmpSol.StudentID,
			&tmpSol.Text, &tmpSol.CreateTime, &tmpSol.File,
			&isApproved, &tmpSol.TeacherEvaluation,
		); err != nil {
			return nil, e.StacktraceError(err)
		}
		tmpSol.Status = makeStatus(isApproved)

		sols = append(sols, &tmpSol)
	}

	return &model.SolutionListForHw{Solutions: sols}, nil
}
