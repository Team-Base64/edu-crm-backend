package usecase

import (
	e "main/domain/errors"
	"main/domain/model"
	"strconv"
	"time"
)

func (uc *Usecase) CreateHomework(teacherID int, newHw *model.HomeworkCreate) (*model.Homework, error) {
	if err := uc.store.CheckClassExistence(newHw.ClassID); err != nil {
		return nil, e.StacktraceError(err)
	}

	createTime := time.Now()
	id, err := uc.store.AddHomework(teacherID, createTime, newHw)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	if err = uc.chatService.BroadcastMsg(genHomeworkMsg(newHw)); err != nil {
		return nil, e.StacktraceError(err, uc.store.DeleteHomework(id))
	}

	tasks, err := uc.store.GetTasksByHomeworkID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	res := &model.Homework{
		ID:           id,
		Title:        newHw.Title,
		Description:  newHw.Description,
		DeadlineTime: newHw.DeadlineTime,
		CreateTime:   createTime,
		Tasks:        tasks,
	}
	return res, nil
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

func genHomeworkMsg(hw *model.HomeworkCreate) *model.ClassBroadcastMessage {
	msg := model.ClassBroadcastMessage{
		ClassID:     hw.ClassID,
		Title:       "Внимание! Выдано домашнее задание: " + hw.Title,
		Description: hw.Description,
	}

	for id, task := range hw.Tasks {
		msg.Description += "\n" + "Задание №" + strconv.Itoa(id) + "\n" + task.Description
		msg.Attaches = append(msg.Attaches, task.Attach)
	}

	msg.Description += "\n" + "Срок выполнения: " + hw.DeadlineTime.String()

	return &msg
}
