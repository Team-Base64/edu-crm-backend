package usecase

import (
	e "main/domain/errors"
	"main/domain/model"
	"time"
)

func (uc *Usecase) CreateHomework(teacherID int, newHw *model.HomeworkCreate) (res *model.Homework, err error) {
	if err = uc.store.CheckClassExistence(newHw.ClassID); err != nil {
		return nil, e.StacktraceError(err)
	}

	createTime := time.Now()
	id, err := uc.store.AddHomework(teacherID, createTime, newHw)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer func() {
		if res == nil {
			uc.store.DeleteHomework(id)
		}
	}()

	bcMsg := model.ClassBroadcastMessage{
		ClassID:     newHw.ClassID,
		Title:       "Внимание! Выдано домашнее задание: " + newHw.Title,
		Description: newHw.Description + "\n" + "Срок выполнения: " + newHw.DeadlineTime.String(),
		Attaches:    []string{newHw.Tasks[0].Attach},
	}
	if err = uc.chatService.BroadcastMsg(&bcMsg); err != nil {
		return nil, e.StacktraceError(err, uc.store.DeleteHomework(id))
	}

	res = &model.Homework{
		ID:           id,
		Title:        newHw.Title,
		Description:  newHw.Description,
		DeadlineTime: newHw.DeadlineTime,
		CreateTime:   createTime,
		File:         newHw.Tasks[0].Attach,
	}
	return
}

func (uc *Usecase) GetHomeworkByID(id int) (*model.HomeworkByID, error) {
	hw, err := uc.store.GetHomeworkByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return hw, nil
}

func (uc *Usecase) GetHomeworksByClassID(classID int) (*model.HomeworkList, error) {
	if err := uc.store.CheckClassExistence(classID); err != nil {
		return nil, e.StacktraceError(err)
	}

	hws, err := uc.store.GetHomeworksByClassID(classID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return hws, nil
}
