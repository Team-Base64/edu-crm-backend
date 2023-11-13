package usecase

import (
	e "main/domain/errors"
	"main/domain/model"
	"math/rand"
)

func (uc *Usecase) CreateClass(teacherID int, newClass *model.ClassCreate) (*model.ClassInfo, error) {
	for i := range uc.bufToken {
		uc.bufToken[i] = uc.letters[rand.Intn(len(uc.letters))]
	}
	inviteToken := string(uc.bufToken)

	id, err := uc.store.AddClass(teacherID, inviteToken, newClass)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	res := model.ClassInfo{
		ID:          id,
		Title:       newClass.Title,
		Description: newClass.Description,
		InviteToken: inviteToken,
	}

	return &res, nil
}

func (uc *Usecase) GetClassesByTeacherID(id int) (*model.ClassInfoList, error) {
	classes, err := uc.store.GetClassesByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return classes, nil
}

func (uc *Usecase) GetClassByID(id int) (*model.ClassInfo, error) {
	class, err := uc.store.GetClassByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return class, nil
}