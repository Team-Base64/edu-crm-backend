package usecase

import (
	e "main/domain/errors"
	"main/domain/model"
	"math/rand"
	"strconv"
	"time"

	"github.com/signintech/gopdf"
)

func (uc *Usecase) CreateHomework(newHw *model.HomeworkCreate) (*model.Homework, error) {
	if err := uc.store.CheckClassExistence(newHw.ClassID); err != nil {
		return nil, e.StacktraceError(err)
	}
	createTime := time.Now()
	id, err := uc.store.AddHomework(createTime, newHw)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	if err := pdf.AddTTFFont("times", "times.ttf"); err != nil {
		return nil, e.StacktraceError(err)
	}
	if err := pdf.SetFont("times", "", 14); err != nil {
		return nil, e.StacktraceError(err)
	}
	for n, task := range newHw.Tasks {
		pdf.AddPage()
		saveTask := &model.TaskByID{}
		if task.ID > 0 {
			saveTask, err = uc.store.GetTaskByID(task.ID)
			if err != nil {
				return nil, e.StacktraceError(err)
			}
		} else {
			saveTask.Description = task.Description
			saveTask.Attach = task.Attach
		}
		if err := pdf.Cell(nil, "Задание № "+strconv.Itoa(n+1)); err != nil {
			return nil, e.StacktraceError(err)
		}
		pdf.Br(20)
		if err := pdf.Cell(nil, saveTask.Description); err != nil {
			return nil, e.StacktraceError(err)
		}

		if len(saveTask.Attach) != 0 {
			pdf.Br(20)
			if err := pdf.Image(saveTask.Attach[21:], 50, 200, nil); err != nil {
				return nil, e.StacktraceError(err)
			}
		}
	}
	for i := range uc.bufToken {
		uc.bufToken[i] = uc.letters[rand.Intn(len(uc.letters))]
	}
	file := string(uc.bufToken) + ".pdf"
	if err := pdf.WritePdf("/filestorage/homework/" + file); err != nil {
		return nil, e.StacktraceError(err)
	}

	bcMsg := model.ClassBroadcastMessage{
		ClassID:     newHw.ClassID,
		Title:       "Внимание! Выдано домашнее задание: " + newHw.Title,
		Description: newHw.Description + "\n" + "Срок выполнения: " + newHw.DeadlineTime.String(),
		Attaches:    []string{"/filestorage/homework/" + file},
	}
	if err := uc.chatService.BroadcastMsg(&bcMsg); err != nil {
		return nil, e.StacktraceError(err, uc.store.DeleteHomework(id))
	}

	res := model.Homework{
		ID:           id,
		Title:        newHw.Title,
		Description:  newHw.Description,
		DeadlineTime: newHw.DeadlineTime,
		CreateTime:   createTime,
		File:         "/filestorage/homework/" + file,
	}
	return &res, nil
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
