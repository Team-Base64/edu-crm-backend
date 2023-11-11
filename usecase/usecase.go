package usecase

import (
	"main/domain/model"
	"math/rand"
	"strconv"
	"time"

	"github.com/signintech/gopdf"

	ctrl "main/controller"
	e "main/domain/errors"
	rep "main/repository"
)

type UsecaseInterface interface {
	CreateTeacher(params *model.TeacherSignUp) error
	GetTeacherProfile(id int) (*model.TeacherProfile, error)
	GetChatsByTeacherID(id int) (*model.ChatPreviewList, error)
	GetChatByID(id int) (*model.Chat, error)
	GetClassesByTeacherID(id int) (*model.ClassInfoList, error)
	GetClassByID(id int) (*model.ClassInfo, error)
	CreateClass(teacherID int, newClass *model.ClassCreate) (*model.ClassInfo, error)
	GetStudentByID(id int) (*model.StudentByID, error)
	GetStudentsFromClass(classID int) (*model.StudentListFromClass, error)
	GetClassFeed(classID int) (*model.Feed, error)
	CreatePost(classID int, newPost *model.PostCreate) (*model.Post, error)
	GetHomeworksByClassID(classID int) (*model.HomeworkList, error)
	GetHomeworkByID(id int) (*model.HomeworkByID, error)
	CreateHomework(newHw *model.HomeworkCreate) (*model.Homework, error)
	GetSolutionsByClassID(classID int) (*model.SolutionListFromClass, error)
	GetSolutionsByHwID(hwID int) (*model.SolutionListForHw, error)
	GetSolutionByID(id int) (*model.SolutionByID, error)
}

type Usecase struct {
	store       rep.StoreInterface
	letters     []rune
	tokenLen    int
	bufToken    []rune
	chatService ctrl.ChatServiceInterface
}

func NewUsecase(s rep.StoreInterface, lettes string, tokenLen int, cs ctrl.ChatServiceInterface) UsecaseInterface {
	return &Usecase{
		store:       s,
		letters:     []rune(lettes),
		tokenLen:    tokenLen,
		bufToken:    make([]rune, tokenLen),
		chatService: cs,
	}
}

func (uc *Usecase) CreateTeacher(params *model.TeacherSignUp) error {
	return uc.store.AddTeacher(params)
}

func (uc *Usecase) GetTeacherProfile(id int) (*model.TeacherProfile, error) {
	chat, err := uc.store.GetTeacherProfile(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return chat, nil
}

func (uc *Usecase) GetChatsByTeacherID(id int) (*model.ChatPreviewList, error) {
	chat, err := uc.store.GetChatsByTeacherID(id)
	if err != nil {

		return nil, e.StacktraceError(err)
	}
	return chat, nil
}

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

func (uc *Usecase) GetClassesByTeacherID(id int) (*model.ClassInfoList, error) {
	classes, err := uc.store.GetClassesByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return classes, nil
}

func (uc *Usecase) GetClassByID(id int) (*model.ClassInfo, error) {
	class, err := uc.store.GetClassByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return class, nil
}

func (uc *Usecase) CreateClass(teacherID int, newClass *model.ClassCreate) (*model.ClassInfo, error) {
	for i := range uc.bufToken {
		uc.bufToken[i] = uc.letters[rand.Intn(len(uc.letters))]
	}
	inviteToken := string(uc.bufToken)

	id, err := uc.store.AddClass(teacherID, inviteToken, newClass)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	res := model.ClassInfo{
		ID:          id,
		Title:       newClass.Title,
		Description: newClass.Description,
		InviteToken: inviteToken,
	}

	return &res, nil
}

func (uc *Usecase) GetStudentByID(id int) (*model.StudentByID, error) {
	student, err := uc.store.GetStudentByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return student, nil
}

func (uc *Usecase) GetStudentsFromClass(classID int) (*model.StudentListFromClass, error) {
	if err := uc.store.CheckClassExistence(classID); err != nil {
		return nil, e.StacktraceError(err)
	}

	students, err := uc.store.GetStudentsFromClass(classID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return students, nil
}

func (uc *Usecase) GetClassFeed(classID int) (*model.Feed, error) {
	if err := uc.store.CheckClassExistence(classID); err != nil {
		return nil, e.StacktraceError(err)
	}

	feed, err := uc.store.GetClassFeed(classID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return feed, nil
}

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

func (uc *Usecase) GetHomeworkByID(id int) (*model.HomeworkByID, error) {
	hw, err := uc.store.GetHomeworkByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return hw, nil
}

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

func (uc *Usecase) GetSolutionsByClassID(classID int) (*model.SolutionListFromClass, error) {
	if err := uc.store.CheckClassExistence(classID); err != nil {
		return nil, e.StacktraceError(err)
	}

	sols, err := uc.store.GetSolutionsByClassID(classID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return sols, nil
}

func (uc *Usecase) GetSolutionsByHwID(hwID int) (*model.SolutionListForHw, error) {
	if err := uc.store.CheckHomeworkExistence(hwID); err != nil {
		return nil, e.StacktraceError(err)
	}

	sols, err := uc.store.GetSolutionsByHwID(hwID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return sols, nil
}

func (uc *Usecase) GetSolutionByID(id int) (*model.SolutionByID, error) {
	sol, err := uc.store.GetSolutionByID(id)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return sol, nil
}
