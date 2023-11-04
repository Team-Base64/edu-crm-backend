package usecase

import (
	"main/domain/model"

	e "main/domain/errors"
	rep "main/repository"
)

type UsecaseInterface interface {
	CreateTeacher(params *model.TeacherSignUp) error
	GetTeacherProfile(id int) (*model.TeacherProfile, error)
	GetChatsByTeacherID(id int) (*model.ChatsPreview, error)
	GetChatByID(id int) (*model.Chat, error)
}

type Usecase struct {
	store rep.StoreInterface
}

func NewUsecase(us rep.StoreInterface) UsecaseInterface {
	return &Usecase{
		store: us,
	}
}

func (api *Usecase) CreateTeacher(params *model.TeacherSignUp) error {
	return api.store.AddTeacher(params)
}

func (api *Usecase) GetTeacherProfile(id int) (*model.TeacherProfile, error) {
	chat, err := api.store.GetTeacherProfile(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return chat, nil
}

func (api *Usecase) GetChatsByTeacherID(id int) (*model.ChatsPreview, error) {
	chat, err := api.store.GetChatsByTeacherID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return chat, nil
}

func (api *Usecase) GetChatByID(id int) (*model.Chat, error) {
	chat, err := api.store.GetChatByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return chat, nil
}
