package pg

import (
	"database/sql"

	e "main/domain/errors"
	m "main/domain/model"

	"github.com/lib/pq"
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

func (s *PostgreSqlStore) GetSolutionByID(id int) (*m.SolutionByID, error) {
	var sol m.SolutionByID
	var isApproved sql.NullBool

	if err := s.db.QueryRow(
		`SELECT id, homeworkID, studentID, text, createTime, files, isApproved, teacherEvaluation FROM solutions WHERE id = $1;`,
		id,
	).Scan(
		&sol.ID, &sol.HomeworkID, &sol.StudentID,
		&sol.Text, &sol.CreateTime, (*pq.StringArray)(&sol.Files), &isApproved, &sol.TeacherEvaluation,
	); err != nil {
		return nil, e.StacktraceError(err)
	}
	sol.Status = makeStatus(isApproved)

	return &sol, nil
}

func (s *PostgreSqlStore) GetSolutionsByClassID(classID int) ([]m.SolutionFromClass, error) {
	rows, err := s.db.Query(
		`SELECT s.id, s.homeworkID, s.studentID, s.text, s.createTime, s.files, s.isApproved, s.teacherEvaluation
		 FROM solutions s
		 JOIN homeworks h ON s.homeworkID = h.id
		 WHERE h.classID = $1;`,
		classID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	sols := []m.SolutionFromClass{}
	for rows.Next() {
		var tmpSol m.SolutionFromClass
		var isApproved sql.NullBool

		if err := rows.Scan(
			&tmpSol.ID, &tmpSol.HomeworkID, &tmpSol.StudentID,
			&tmpSol.Text, &tmpSol.CreateTime, (*pq.StringArray)(&tmpSol.Files),
			&isApproved, &tmpSol.TeacherEvaluation,
		); err != nil {
			return nil, e.StacktraceError(err)
		}
		tmpSol.Status = makeStatus(isApproved)

		sols = append(sols, tmpSol)
	}

	return sols, nil
}

func (s *PostgreSqlStore) GetSolutionsByHomeworkID(homeworkID int) ([]m.SolutionForHw, error) {
	rows, err := s.db.Query(
		`SELECT id, studentID, text, createTime, files, isApproved, teacherEvaluation FROM solutions WHERE homeworkID = $1;`,
		homeworkID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	sols := []m.SolutionForHw{}
	for rows.Next() {
		var tmpSol m.SolutionForHw
		var isApproved sql.NullBool

		if err := rows.Scan(
			&tmpSol.ID, &tmpSol.StudentID,
			&tmpSol.Text, &tmpSol.CreateTime, (*pq.StringArray)(&tmpSol.Files),
			&isApproved, &tmpSol.TeacherEvaluation,
		); err != nil {
			return nil, e.StacktraceError(err)
		}
		tmpSol.Status = makeStatus(isApproved)

		sols = append(sols, tmpSol)
	}

	return sols, nil
}

func (s *PostgreSqlStore) GetInfoForEvaluationMsgBySolutionID(solutionID int) (*m.SolutionInfoForEvaluationMsg, error) {
	var info m.SolutionInfoForEvaluationMsg

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

func (s *PostgreSqlStore) AddEvaluationForSolution(solutionID int, isApproved bool, evaluation string) error {
	if _, err := s.db.Exec(
		`UPDATE solutions SET isapproved = $1, teacherevaluation = $2 WHERE id = $3;`,
		isApproved, evaluation, solutionID,
	); err != nil {
		return e.StacktraceError(err)
	}

	return nil
}
