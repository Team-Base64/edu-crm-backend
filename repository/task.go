package repository

import (
	e "main/domain/errors"
	"main/domain/model"
)

func (s *Store) AddTask(newTask *model.TaskCreate) (int, error) {
	var id int
	if err := s.db.QueryRow(
		`INSERT INTO tasks (description, attach)
		 VALUES ($1, $2)
		 RETURNING id;`,
		newTask.Description, newTask.Attach,
	).Scan(&id); err != nil {
		return 0, e.StacktraceError(err)
	}

	return int(id), nil
}

func (s *Store) GetTaskByID(id int) (*model.TaskByID, error) {
	var task model.TaskByID
	if err := s.db.QueryRow(
		`SELECT description, attach FROM solutions WHERE id = $1;`,
		id,
	).Scan(
		&task.Description, &task.Attach,
	); err != nil {
		return nil, e.StacktraceError(err)
	}

	return &task, nil
}

func (s *Store) AttachTaskToHomework(hwID int, taskID int) error {
	_, err := s.db.Exec(
		`INSERT INTO homeworks_tasks (homeworkID, taskID) VALUES ($1, $2)`,
		hwID, taskID,
	)
	if err != nil {
		return e.StacktraceError(err)
	}

	return nil
}
