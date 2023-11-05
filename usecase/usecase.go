package usecase

import (
	"main/domain/model"
	"math/rand"

	e "main/domain/errors"
	rep "main/repository"
)

type UsecaseInterface interface {
	CreateTeacher(params *model.TeacherSignUp) error
	GetTeacherProfile(id int) (*model.TeacherProfile, error)
	GetChatsByTeacherID(id int) (*model.ChatPreviewList, error)
	GetChatByID(id int) (*model.Chat, error)
	GetClassesByTeacherID(id int) (*model.ClassesInfo, error)
	GetClassByID(id int) (*model.ClassInfo, error)
	CreateClass(teacherID int, newClass *model.ClassCreate) (*model.ClassCreateResponse, error)
	GetStudentsFromClass(classID int) (*model.StudentsFromClass, error)
	GetClassFeed(classID int) (*model.Feed, error)
}

type Usecase struct {
	store    rep.StoreInterface
	letters  []rune
	tokenLen int
	bufToken []rune
}

func NewUsecase(s rep.StoreInterface, lettes string, tokenLen int) UsecaseInterface {
	return &Usecase{
		store:    s,
		letters:  []rune(lettes),
		tokenLen: tokenLen,
		bufToken: make([]rune, tokenLen),
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

func (api *Usecase) GetClassesByTeacherID(id int) (*model.ClassesInfo, error) {
	classes, err := api.store.GetClassesByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return classes, nil
}

func (api *Usecase) GetClassByID(id int) (*model.ClassInfo, error) {
	class, err := api.store.GetClassByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return class, nil
}

func (api *Usecase) CreateClass(teacherID int, newClass *model.ClassCreate) (*model.ClassCreateResponse, error) {
	for i := range api.bufToken {
		api.bufToken[i] = api.letters[rand.Intn(len(api.letters))]
	}
	inviteToken := string(api.bufToken)

	id, err := api.store.AddClass(teacherID, inviteToken, newClass)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	res := model.ClassCreateResponse{
		ID:          id,
		InviteToken: inviteToken,
	}

	return &res, nil
}

func (api *Usecase) GetStudentsFromClass(classID int) (*model.StudentsFromClass, error) {
	students, err := api.store.GetStudentsFromClass(classID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return students, nil
}

func (api *Usecase) GetClassFeed(classID int) (*model.Feed, error) {
	feed, err := api.store.GetClassFeed(classID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return feed, nil
}
