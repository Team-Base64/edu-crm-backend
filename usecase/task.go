package usecase

import (
	e "main/domain/errors"
	"main/domain/model"
)

func (uc *Usecase) GetTasksByTeacher(teacherID int) (*model.TaskListByTeacherID, error) {
	tasks, err := uc.store.GetTasksByTeacher(teacherID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	return tasks, nil
}
