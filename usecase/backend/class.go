package backend

import (
	e "main/domain/errors"
	m "main/domain/model"
)

func (uc *BackendUsecase) CreateClass(teacherID int, newClass *m.ClassCreate) (*m.ClassInfo, error) {
	inviteToken := uc.genRandomToken()

	id, err := uc.dataStore.AddClass(teacherID, inviteToken, newClass)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	res := m.ClassInfo{
		ID:          id,
		Title:       newClass.Title,
		Description: newClass.Description,
		InviteToken: inviteToken,
	}

	return &res, nil
}

func (uc *BackendUsecase) GetClassesByTeacherID(id int) ([]m.ClassInfo, error) {
	classes, err := uc.dataStore.GetClassesByTeacherID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return classes, nil
}

func (uc *BackendUsecase) GetClassByID(id int) (*m.ClassInfo, error) {
	class, err := uc.dataStore.GetClassByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return class, nil
}
