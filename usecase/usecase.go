package usecase

import (
	"main/domain/model"
	rep "main/repository"

	chat "main/microservices/chatServer/gen_files"

	"github.com/google/uuid"
)

type UsecaseInterface interface {
	CreateTeacher(params *model.TeacherDB) error
	GetTeacher(id int) (*model.TeacherDB, error)
	ChangeTeacher(params *model.TeacherDB) error
	GetChatsByTeacherID(id int) (model.Chats, error)
	AddStudent(params *model.CreateStudentDB) error
	GetChatByID(id int) (model.Chat, error)
}

type Usecase struct {
	chatManager chat.BotChatClient
	store       rep.StoreInterface
}

func NewUsecase(us rep.StoreInterface, cm chat.BotChatClient) UsecaseInterface {
	return &Usecase{
		chatManager: cm,
		store:       us,
	}
}

func (api *Usecase) CreateTeacher(params *model.TeacherDB) error {
	return api.store.AddTeacher(params)
}

func (api *Usecase) GetTeacher(id int) (*model.TeacherDB, error) {
	return api.store.GetTeacher(id)
}

func (api *Usecase) ChangeTeacher(params *model.TeacherDB) error {
	return api.store.UpdateTeacher(params)
}

func (api *Usecase) AddStudent(params *model.CreateStudentDB) error {
	newUUID := uuid.New()
	api.store.CreateChat(&model.ChatDB{TeacherID: 1, StudentHash: newUUID.String()})
	return api.store.AddStudent(&model.StudentDB{InviteHash: newUUID.String(), Name: params.Name})
}

func (api *Usecase) GetChatsByTeacherID(id int) (model.Chats, error) {
	chats, err := api.store.GetChatsByID(id)
	return *chats, err
}

func (api *Usecase) GetChatByID(id int) (model.Chat, error) {
	chat, err := api.store.GetChatFromDB(id)
	return *chat, err
}
