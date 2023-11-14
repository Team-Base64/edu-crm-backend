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
