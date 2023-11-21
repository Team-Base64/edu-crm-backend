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
		`SELECT id, homeworkID, studentID, text, createTime, file, isApproved, teacherEvaluation FROM solutions WHERE id = $1;`,
		id,
	).Scan(
		&sol.ID, &sol.HomeworkID, &sol.StudentID,
		&sol.Text, &sol.CreateTime, &sol.File, &isApproved, &sol.TeacherEvaluation,
	); err != nil {
		return nil, e.StacktraceError(err)
	}
	sol.Status = makeStatus(isApproved)

	return &sol, nil
}

func (s *Store) GetSolutionsByClassID(classID int) ([]model.SolutionFromClass, error) {
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

	sols := []model.SolutionFromClass{}
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

		sols = append(sols, tmpSol)
	}

	return sols, nil
}

func (s *Store) GetSolutionsByHomeworkID(homeworkID int) ([]model.SolutionForHw, error) {
	rows, err := s.db.Query(
		`SELECT id, studentID, text, createTime, file, isApproved, teacherEvaluation FROM solutions WHERE homeworkID = $1;`,
		homeworkID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	sols := []model.SolutionForHw{}
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

		sols = append(sols, tmpSol)
	}

	return sols, nil
}

func (s *Store) GetInfoForEvaluationMsgBySolutionID(solutionID int) (*model.SolutionInfoForEvaluationMsg, error) {
	var info model.SolutionInfoForEvaluationMsg

	if err := s.db.QueryRow(
		`SELECT h.title, s.createTime FROM homeworks h
		 JOIN solutions s ON h.id = s.homeworkID
		 WHERE s.id = $1;`,
		solutionID,
	).Scan(&info.HomeworkTitle, &info.SolutionCreateTime); err != nil {
		return nil, e.StacktraceError(err)
	}

	return &info, nil
}

func (s *Store) AddEvaluationForSolution(solutionID int, isApproved bool, evaluation string) error {
	if _, err := s.db.Exec(
		`UPDATE solutions SET isapproved = $1, teacherevaluation = $2 WHERE id = $3;`,
		isApproved, evaluation, solutionID,
	); err != nil {
		return e.StacktraceError(err)
	}

	return nil
}
