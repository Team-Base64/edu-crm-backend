package usecase

import (
	e "main/domain/errors"
	"main/domain/model"
	"main/domain/utils"
	"strconv"
	"time"

	"log"
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

	msg := uc.genHomeworkMsg(newHw)
	if err = uc.ctrlService.BroadcastMsg(&msg); err != nil {
		return nil, e.StacktraceError(err, uc.store.DeleteHomework(id))
	}

	res := &model.Homework{
		ID:           id,
		Title:        newHw.Title,
		Description:  newHw.Description,
		DeadlineTime: newHw.DeadlineTime,
		CreateTime:   createTime,
		Tasks:        newHw.Tasks,
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
	return &model.HomeworkList{Homeworks: hws}, nil
}

func (uc Usecase) genHomeworkMsg(hw *model.HomeworkCreate) model.ClassBroadcastMessage {
	log.Println(hw.DeadlineTime)
	msg := model.ClassBroadcastMessage{
		ClassID: hw.ClassID,
		Title:   "Внимание! Выдано домашнее задание: " + hw.Title,
		Description: "Задач в д/з: " + strconv.Itoa(len(hw.Tasks)) + "\n" +
			"Срок выполнения: " + utils.TimeToString(hw.DeadlineTime),
	}

	return msg
}
