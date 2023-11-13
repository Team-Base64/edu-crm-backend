package repository

import (
	e "main/domain/errors"
	"main/domain/model"
)

func (s *Store) AddTask(teacherID int, newTask *model.TaskCreate) (int, error) {
	var id int
	if err := s.db.QueryRow(
		`INSERT INTO tasks (teacherID, description, attach)
		 VALUES ($1, $2, $3)
		 RETURNING id;`,
		teacherID, newTask.Description, newTask.Attach,
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

func (s *Store) GetTasksByTeacher(teacherID int) (*model.TaskListByTeacherID, error) {
	rows, err := s.db.Query(
		`SELECT id, description, attach FROM solutions WHERE teacherID = $1;`,
		teacherID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	tasks := []model.Task{}
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Description, &task.Attach)
		if err != nil {
			return nil, e.StacktraceError(err)
		}

		tasks = append(tasks, task)
	}

	return &model.TaskListByTeacherID{Tasks: tasks}, nil
}
