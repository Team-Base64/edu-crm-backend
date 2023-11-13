package usecase

import (
	e "main/domain/errors"
	"main/domain/model"
)

func (uc *Usecase) GetStudentByID(id int) (*model.StudentByID, error) {
	student, err := uc.store.GetStudentByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return student, nil
}

func (uc *Usecase) GetStudentsFromClass(classID int) (*model.StudentListFromClass, error) {
	if err := uc.store.CheckClassExistence(classID); err != nil {
		return nil, e.StacktraceError(err)
	}

	students, err := uc.store.GetStudentsFromClass(classID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return students, nil
}
