package repository

import (
	e "main/domain/errors"
	"main/domain/model"
	"time"
)

func (s *Store) CheckHomeworkExistence(id int) error {
	var tmp int
	if err := s.db.QueryRow(
		`SELECT 1 FROM homeworks WHERE id = $1;`,
		id,
	).Scan(&tmp); err != nil {
		return e.StacktraceError(err)
	}
	return nil
}

func (s *Store) AddHomework(teacherID int, createTime time.Time, newHw *model.HomeworkCreate) (int, error) {
	var hwID int
	if err := s.db.QueryRow(
		`INSERT INTO homeworks (classID, title, description, createTime, deadlineTime)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING id;`,
		newHw.ClassID, newHw.Title, newHw.Description, createTime, newHw.DeadlineTime,
	).Scan(&hwID); err != nil {
		return 0, e.StacktraceError(err)
	}

	for rank, taskID := range newHw.Tasks {
		if err := s.AttachTaskToHomework(hwID, taskID, rank); err != nil {
			return 0, e.StacktraceError(err, s.DeleteHomework(hwID))
		}
	}

	return hwID, nil
}

func (s *Store) DeleteHomework(id int) error {
	_, err := s.db.Exec(
		`DELETE FROM homeworks WHERE id = $1;`,
		id,
	)

	if err != nil {
		return e.StacktraceError(err)
	}

	return nil
}

func (s *Store) GetHomeworkByID(id int) (*model.HomeworkByID, error) {
	var hw model.HomeworkByID
	if err := s.db.QueryRow(
		`SELECT classID, title, description, createTime, deadlineTime
		 FROM homeworks
		 WHERE id = $1;`,
		id,
	).Scan(
		&hw.ClassID, &hw.Title, &hw.Description,
		&hw.CreateTime, &hw.DeadlineTime,
	); err != nil {
		return nil, e.StacktraceError(err)
	}

	tasks, err := s.GetTasksIDByHomeworkID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	hw.Tasks = tasks

	return &hw, nil
}

func (s *Store) GetHomeworksByClassID(classID int) ([]model.Homework, error) {
	rows, err := s.db.Query(
		`SELECT id, title, description, createTime, deadlineTime
		 FROM homeworks
		 WHERE classID = $1;`,
		classID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	var hws []model.Homework
	for rows.Next() {
		var tmpHw model.Homework

		if err := rows.Scan(
			&tmpHw.ID, &tmpHw.Title, &tmpHw.Description,
			&tmpHw.CreateTime, &tmpHw.DeadlineTime,
		); err != nil {
			return nil, e.StacktraceError(err)
		}

		tasks, err := s.GetTasksIDByHomeworkID(tmpHw.ID)
		if err != nil {
			return nil, e.StacktraceError(err)
		}
		tmpHw.Tasks = tasks

		hws = append(hws, tmpHw)
	}

	return hws, nil
}
