package pg

import (
	e "main/domain/errors"
	m "main/domain/model"

	"github.com/lib/pq"
)

func (s *PostgreSqlStore) AddTask(teacherID int, newTask *m.TaskCreate) (int, error) {
	var id int
	if newTask.Attaches == nil {
		newTask.Attaches = []string{}
	}

	if err := s.db.QueryRow(
		`INSERT INTO tasks (teacherID, description, attaches)
		 VALUES ($1, $2, $3)
		 RETURNING id;`,
		teacherID, newTask.Description, (*pq.StringArray)(&newTask.Attaches),
	).Scan(&id); err != nil {
		return 0, e.StacktraceError(err)
	}

	return int(id), nil
}

func (s *PostgreSqlStore) GetTaskByID(id int) (*m.TaskByID, error) {
	var task m.TaskByID
	if err := s.db.QueryRow(
		`SELECT description, attaches FROM tasks WHERE id = $1;`,
		id,
	).Scan(
		&task.Description, (*pq.StringArray)(&task.Attaches),
	); err != nil {
		return nil, e.StacktraceError(err)
	}

	return &task, nil
}

func (s *PostgreSqlStore) GetTasksByTeacherID(teacherID int) ([]m.Task, error) {
	rows, err := s.db.Query(
		`SELECT id, description, attaches FROM tasks WHERE teacherID = $1;`,
		teacherID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	tasks := []m.Task{}
	for rows.Next() {
		var task m.Task
		err := rows.Scan(&task.ID, &task.Description, (*pq.StringArray)(&task.Attaches))
		if err != nil {
			return nil, e.StacktraceError(err)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *PostgreSqlStore) GetTasksByHomeworkID(homeworkID int) ([]m.Task, error) {
	rows, err := s.db.Query(
		`SELECT t.id, t.description, t.attaches
		 FROM tasks t
		 JOIN homeworks_tasks ht ON t.id = ht.taskID
		 WHERE ht.homeworkID = $1
		 ORDER BY ht.rank;`,
		homeworkID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	tasks := []m.Task{}
	for rows.Next() {
		var task m.Task
		err := rows.Scan(&task.ID, &task.Description, (*pq.StringArray)(&task.Attaches))
		if err != nil {
			return nil, e.StacktraceError(err)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *PostgreSqlStore) GetTasksIDByHomeworkID(homeworkID int) ([]int, error) {
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

	tasks := []int{}
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

func (s *PostgreSqlStore) AttachTaskToHomework(hwID int, taskID int, taskRank int) error {
	_, err := s.db.Exec(
		`INSERT INTO homeworks_tasks (homeworkID, taskID, rank) VALUES ($1, $2, $3)`,
		hwID, taskID, taskRank,
	)
	if err != nil {
		return e.StacktraceError(err)
	}

	return nil
}
