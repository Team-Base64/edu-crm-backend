package backend

import (
	"log"
	"strconv"
	"time"

	e "main/domain/errors"
	m "main/domain/model"
	u "main/domain/utils"
)

func (uc *BackendUsecase) CreateHomework(teacherID int, newHw *m.HomeworkCreate) (*m.Homework, error) {
	if err := uc.dataStore.CheckClassExistence(newHw.ClassID); err != nil {
		return nil, e.StacktraceError(err)
	}

	createTime := time.Now()
	id, err := uc.dataStore.AddHomework(teacherID, createTime, newHw)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	msg := uc.genHomeworkMsg(newHw)
	if err = uc.chat.BroadcastMsg(&msg); err != nil {
		return nil, e.StacktraceError(err, uc.dataStore.DeleteHomework(id))
	}

	res := &m.Homework{
		ID:           id,
		Title:        newHw.Title,
		Description:  newHw.Description,
		DeadlineTime: newHw.DeadlineTime,
		CreateTime:   createTime,
		Tasks:        newHw.Tasks,
	}
	return res, nil
}

func (uc *BackendUsecase) GetHomeworkByID(id int) (*m.HomeworkByID, error) {
	hw, err := uc.dataStore.GetHomeworkByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return hw, nil
}

func (uc *BackendUsecase) GetHomeworksByClassID(classID int) (*m.HomeworkList, error) {
	if err := uc.dataStore.CheckClassExistence(classID); err != nil {
		return nil, e.StacktraceError(err)
	}

	hws, err := uc.dataStore.GetHomeworksByClassID(classID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return &m.HomeworkList{Homeworks: hws}, nil
}

func (uc BackendUsecase) genHomeworkMsg(hw *m.HomeworkCreate) m.ClassBroadcastMessage {
	log.Println(hw.DeadlineTime)
	msg := m.ClassBroadcastMessage{
		ClassID: hw.ClassID,
		Title:   "Внимание! Выдано домашнее задание: " + hw.Title,
		Description: "Задач в д/з: " + strconv.Itoa(len(hw.Tasks)) + "\n" +
			"Срок выполнения: " + u.TimeToString(hw.DeadlineTime),
	}

	return msg
}
