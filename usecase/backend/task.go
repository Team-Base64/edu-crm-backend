package backend

import (
	e "main/domain/errors"
	m "main/domain/model"
)

func (uc *BackendUsecase) CreateTask(teacherID int, newTask *m.TaskCreate) (*m.TaskCreateResponse, error) {
	id, err := uc.dataStore.AddTask(teacherID, newTask)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	return &m.TaskCreateResponse{ID: id}, nil
}

func (uc *BackendUsecase) GetTaskByID(id int) (*m.TaskByID, error) {
	task, err := uc.dataStore.GetTaskByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	return task, nil
}

func (uc *BackendUsecase) GetTasksByTeacherID(teacherID int) ([]m.Task, error) {
	tasks, err := uc.dataStore.GetTasksByTeacherID(teacherID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	return tasks, nil
}
