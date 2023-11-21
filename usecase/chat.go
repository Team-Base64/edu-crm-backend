package usecase

import (
	e "main/domain/errors"
	"main/domain/model"
)

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

func (uc *Usecase) GetChatsByTeacherID(id int) ([]model.ChatPreview, error) {
	chats, err := uc.store.GetChatsByTeacherID(id)
	if err != nil {

		return nil, e.StacktraceError(err)
	}
	return chats, nil
}
