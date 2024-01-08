package usecase

import (
	"main/domain/model"
	"math/rand"

	d "main/delivery"
	rep "main/repository"
)

type UsecaseInterface interface {
	// TEACHER
	SignUpTeacher(req *model.TeacherSignUp) error
	CreateTeacher(params *model.TeacherSignUp) error
	GetTeacherProfile(id int) (*model.TeacherProfile, error)
	GetTeacherProfileByLogin(login string) (*model.TeacherDB, error)
	// SESSION
	CreateSession(teacherLogin string) (*model.Session, error)
	CheckSession(in string) (string, error)
	DeleteSession(in string) error
	// CHAT
	GetChatByID(id int) (*model.Chat, error)
	GetChatsByTeacherID(id int) ([]model.ChatPreview, error)
	ReadChatByID(id int, teacherID int) error
	// CLASS
	CreateClass(teacherID int, newClass *model.ClassCreate) (*model.ClassInfo, error)
	GetClassByID(id int) (*model.ClassInfo, error)
	GetClassesByTeacherID(id int) ([]model.ClassInfo, error)
	// STUDENT
	GetStudentByID(id int) (*model.StudentByID, error)
	GetStudentsFromClass(classID int) ([]model.StudentFromClass, error)
	// FEED
	CreatePost(classID int, newPost *model.PostCreate) (*model.Post, error)
	GetClassPosts(classID int) ([]model.Post, error)
	// HOMEWORK
	CreateHomework(teacherID int, newHw *model.HomeworkCreate) (*model.Homework, error)
	GetHomeworkByID(id int) (*model.HomeworkByID, error)
	GetHomeworksByClassID(classID int) (*model.HomeworkList, error)
	// TASKS
	CreateTask(teacherID int, newTask *model.TaskCreate) (*model.TaskCreateResponse, error)
	GetTaskByID(id int) (*model.TaskByID, error)
	GetTasksByTeacherID(teacherID int) ([]model.Task, error)
	// SOLUTION
	GetSolutionByID(id int) (*model.SolutionByID, error)
	GetSolutionsByClassID(classID int) ([]model.SolutionFromClass, error)
	GetSolutionsByHomeworkID(homeworkID int) ([]model.SolutionForHw, error)
	EvaluateSolutionbyID(solutionID int, evaluation *model.SolutionEvaluation) error
	// CALENDAR
	// CreateCalendar(teacherID int) (*model.CalendarParams, error)
	GetCalendar(teacherID int) (*model.CalendarParams, error)
	// CreateCalendarEvent(req *model.CalendarEvent, teacherID int) error
	GetCalendarEvents(teacherID int) (model.CalendarEvents, error)
	// DeleteCalendarEvent(teacherID int, eventID string) error
	// UpdateCalendarEvent(req *model.CalendarEvent, teacherID int) error
}

type Usecase struct {
	store           rep.StoreInterface
	letters         []rune
	tokenLen        int
	bufToken        []rune
	chat            d.ChatInterface
	calendar        d.CalendarInterface
	tokenFile       string
	credentialsFile string
}

func NewUsecase(
	s rep.StoreInterface,
	lettes string,
	tokenLen int,
	chat d.ChatInterface,
	calendar d.CalendarInterface,
	tok string,
	cred string,
) UsecaseInterface {
	return &Usecase{
		store:           s,
		letters:         []rune(lettes),
		tokenLen:        tokenLen,
		bufToken:        make([]rune, tokenLen),
		chat:            chat,
		calendar:        calendar,
		tokenFile:       tok,
		credentialsFile: cred,
	}
}

func (uc Usecase) genRandomToken() string {
	for i := range uc.bufToken {
		uc.bufToken[i] = uc.letters[rand.Intn(len(uc.letters))]
	}
	return string(uc.bufToken)
}
