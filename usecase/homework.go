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

	// file, err := uc.genHomeworkPDF(newHw.Tasks)
	// if err != nil {
	// 	return nil, e.StacktraceError(err)
	// }

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

// func (uc Usecase) genHomeworkPDF(tasks []*model.Task) (string, error) {
// 	pdf := gopdf.GoPdf{}
// 	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

// 	if err := pdf.AddTTFFont("times", "times.ttf"); err != nil {
// 		return "", e.StacktraceError(err)
// 	}
// 	if err := pdf.SetFont("times", "", 14); err != nil {
// 		return "", e.StacktraceError(err)
// 	}

// 	for n, task := range tasks {
// 		pdf.AddPage()
// 		saveTask := &model.TaskByID{}

// 		if task.ID > 0 {
// 			storeTask, err := uc.store.GetTaskByID(task.ID)
// 			if err != nil {
// 				return "", e.StacktraceError(err)
// 			}
// 			saveTask = storeTask
// 		} else {
// 			saveTask.Description = task.Description
// 			saveTask.Attach = task.Attach
// 		}

// 		if err := pdf.Cell(nil, "Задание № "+strconv.Itoa(n+1)); err != nil {
// 			return "", e.StacktraceError(err)
// 		}

// 		pdf.Br(20)
// 		if err := pdf.Cell(nil, saveTask.Description); err != nil {
// 			return "", e.StacktraceError(err)
// 		}

// 		if len(saveTask.Attach) != 0 {
// 			if err := pdf.Image(saveTask.Attach[21:], 50, 200, nil); err != nil {
// 				return "", e.StacktraceError(err)
// 			}
// 		}
// 	}

// 	file := "/filestorage/homework/" + uuid.New().String() + ".pdf"
// 	if err := pdf.WritePdf(file); err != nil {
// 		return "", e.StacktraceError(err)
// 	}
// 	return file, nil
// }
