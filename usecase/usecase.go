package usecase

import (
	"main/domain/model"

	e "main/domain/errors"
	rep "main/repository"
)

type UsecaseInterface interface {
	CreateTeacher(params *model.TeacherSignUp) error
	GetTeacherProfile(id int) (*model.TeacherProfile, error)
	GetChatsByTeacherID(id int) (*model.ChatPreviewList, error)
	GetChatByID(id int) (*model.Chat, error)
	GetClassesByTeacherID(id int) (*model.Classes, error)
	GetClassByID(id int) (*model.Class, error)
}

type Usecase struct {
	store rep.StoreInterface
}

func NewUsecase(s rep.StoreInterface) UsecaseInterface {
	return &Usecase{
		store: s,
	}
}

func (uc *Usecase) CreateTeacher(params *model.TeacherSignUp) error {
	return uc.store.AddTeacher(params)
}

func (uc *Usecase) GetTeacherProfile(id int) (*model.TeacherProfile, error) {
	chat, err := uc.store.GetTeacherProfile(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return chat, nil
}

func (uc *Usecase) GetChatsByTeacherID(id int) (*model.ChatPreviewList, error) {
	chat, err := uc.store.GetChatsByTeacherID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return chat, nil
}

func (uc *Usecase) GetChatByID(id int) (*model.Chat, error) {
	if err := uc.store.CheckChatExistence(id); err != nil {
		return nil, e.StacktraceError(err)
	}

	chat, err := uc.store.GetChatByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return chat, nil
}

func (api *Usecase) GetClassesByTeacherID(id int) (*model.Classes, error) {
	classes, err := api.store.GetClassesByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return classes, nil
}

func (api *Usecase) GetClassByID(id int) (*model.Class, error) {
	class, err := api.store.GetClassByID(id)
	if err != nil {
		log.Println("store: ", err)
		return nil, err
	}
	return class, nil
}
