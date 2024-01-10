package backend

import (
	e "main/domain/errors"
	m "main/domain/model"
)

func (uc *BackendUsecase) GetStudentByID(id int) (*m.StudentByID, error) {
	student, err := uc.dataStore.GetStudentByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return student, nil
}

func (uc *BackendUsecase) GetStudentsFromClass(classID int) ([]m.StudentFromClass, error) {
	if err := uc.dataStore.CheckClassExistence(classID); err != nil {
		return nil, e.StacktraceError(err)
	}

	students, err := uc.dataStore.GetStudentsFromClass(classID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return students, nil
}
