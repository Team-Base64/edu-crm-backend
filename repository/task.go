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
		`SELECT description, attach FROM tasks WHERE id = $1;`,
		id,
	).Scan(
		&task.Description, &task.Attach,
	); err != nil {
		return nil, e.StacktraceError(err)
	}

	return &task, nil
}

func (s *Store) GetTasksByTeacherID(teacherID int) ([]*model.Task, error) {
	rows, err := s.db.Query(
		`SELECT id, description, attach FROM tasks WHERE teacherID = $1;`,
		teacherID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	tasks := []*model.Task{}
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Description, &task.Attach)
		if err != nil {
			return nil, e.StacktraceError(err)
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (s *Store) GetTasksByHomeworkID(homeworkID int) ([]*model.Task, error) {
	rows, err := s.db.Query(
		`SELECT t.id, t.description, t.attach
		 FROM tasks
		 JOIN homeworks_tasks ht ON t.id = ht.taskID
		 WHERE ht.homeworkID = $1
		 ORDER BY ht.rank;`,
		homeworkID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	tasks := []*model.Task{}
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Description, &task.Attach)
		if err != nil {
			return nil, e.StacktraceError(err)
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (s *Store) GetTasksIDByHomeworkID(homeworkID int) ([]int, error) {
	rows, err := s.db.Query(
		`SELECT taskID FROM homeworks_tasks
		 WHERE homeworkID = $1
		 ORDER BY rank;`,
		homeworkID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	var tasks []int
	for rows.Next() {
		var taskID int
		err := rows.Scan(&taskID)
		if err != nil {
			return nil, e.StacktraceError(err)
		}

		tasks = append(tasks, taskID)
	}

	return tasks, nil
}

func (s *Store) AttachTaskToHomework(hwID int, taskID int, taskRank int) error {
	_, err := s.db.Exec(
		`INSERT INTO homeworks_tasks (homeworkID, taskID) VALUES ($1, $2, $3)`,
		hwID, taskID, taskRank,
	)
	if err != nil {
		return e.StacktraceError(err)
	}

	return nil
}
