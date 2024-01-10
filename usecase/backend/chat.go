package backend

import (
	e "main/domain/errors"
	m "main/domain/model"
)

func (uc *BackendUsecase) GetChatByID(id int) (*m.Chat, error) {
	if err := uc.dataStore.CheckChatExistence(id); err != nil {
		return nil, e.StacktraceError(err)
	}

	chat, err := uc.dataStore.GetChatByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return chat, nil
}

func (uc *BackendUsecase) GetChatsByTeacherID(id int) ([]m.ChatPreview, error) {
	chats, err := uc.dataStore.GetChatsByTeacherID(id)
	if err != nil {

		return nil, e.StacktraceError(err)
	}
	return chats, nil
}

func (uc *BackendUsecase) ReadChatByID(id int, teacherID int) error {
	err := uc.dataStore.ReadChatByID(id, teacherID)
	if err != nil {
		return e.StacktraceError(err)
	}
	return nil
}
