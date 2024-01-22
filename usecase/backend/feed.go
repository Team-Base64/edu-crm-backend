package backend

import (
	"time"

	e "main/domain/errors"
	m "main/domain/model"
)

func (uc *BackendUsecase) CreatePost(classID int, newPost *m.PostCreate) (*m.Post, error) {
	if err := uc.dataStore.CheckClassExistence(classID); err != nil {
		return nil, e.StacktraceError(err)
	}
	createTime := time.Now()
	id, err := uc.dataStore.AddPost(classID, createTime, newPost)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	bcMsg := m.ClassBroadcastMessage{
		ClassID:     classID,
		Title:       "Внимание! Сообщение от преподавателя.",
		Description: newPost.Text,
		Attaches:    newPost.Attaches,
	}
	if err := uc.chat.BroadcastMsg(&bcMsg); err != nil {
		return nil, e.StacktraceError(err, uc.dataStore.DeletePost(id))
	}

	res := m.Post{
		ID:         id,
		Text:       newPost.Text,
		Attaches:   newPost.Attaches,
		CreateTime: createTime,
	}
	return &res, nil
}

func (uc *BackendUsecase) GetClassPosts(classID int) ([]m.Post, error) {
	if err := uc.dataStore.CheckClassExistence(classID); err != nil {
		return nil, e.StacktraceError(err)
	}

	posts, err := uc.dataStore.GetClassPosts(classID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return posts, nil
}
