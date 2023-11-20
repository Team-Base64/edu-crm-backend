package usecase

import (
	e "main/domain/errors"
	"main/domain/model"
	"time"
)

func (uc *Usecase) CreatePost(classID int, newPost *model.PostCreate) (*model.Post, error) {
	if err := uc.store.CheckClassExistence(classID); err != nil {
		return nil, e.StacktraceError(err)
	}
	createTime := time.Now()
	id, err := uc.store.AddPost(classID, createTime, newPost)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	bcMsg := model.ClassBroadcastMessage{
		ClassID:     classID,
		Title:       "Внимание! Сообщение от преподавателя.",
		Description: newPost.Text,
		Attaches:    newPost.Attaches,
	}
	if err := uc.chatService.BroadcastMsg(&bcMsg); err != nil {
		return nil, e.StacktraceError(err, uc.store.DeletePost(id))
	}

	res := model.Post{
		ID:         id,
		Text:       newPost.Text,
		Attaches:   newPost.Attaches,
		CreateTime: createTime,
	}
	return &res, nil
}

func (uc *Usecase) GetClassPosts(classID int) ([]model.Post, error) {
	if err := uc.store.CheckClassExistence(classID); err != nil {
		return nil, e.StacktraceError(err)
	}

	posts, err := uc.store.GetClassPosts(classID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return posts, nil
}
