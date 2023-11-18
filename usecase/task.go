package usecase

import (
	e "main/domain/errors"
	"main/domain/model"
)

func (uc *Usecase) GetTasksByTeacherID(teacherID int) ([]*model.Task, error) {
	tasks, err := uc.store.GetTasksByTeacherID(teacherID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	return tasks, nil
}

func (uc *Usecase) CreateTask(teacherID int, newTask *model.TaskCreate) (*model.TaskCreateResponse, error) {
	id, err := uc.store.AddTask(teacherID, newTask)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	return &model.TaskCreateResponse{ID: id}, nil
}
