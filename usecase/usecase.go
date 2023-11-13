package usecase

import (
	"main/domain/model"

	ctrl "main/controller"
	rep "main/repository"
)

type UsecaseInterface interface {
	// TEACHER
	CreateTeacher(params *model.TeacherSignUp) error
	GetTeacherProfile(id int) (*model.TeacherProfile, error)
	// CHAT
	GetChatByID(id int) (*model.Chat, error)
	GetChatsByTeacherID(id int) (*model.ChatPreviewList, error)
	// CLASS
	CreateClass(teacherID int, newClass *model.ClassCreate) (*model.ClassInfo, error)
	GetClassByID(id int) (*model.ClassInfo, error)
	GetClassesByTeacherID(id int) (*model.ClassInfoList, error)
	// STUDENT
	GetStudentByID(id int) (*model.StudentByID, error)
	GetStudentsFromClass(classID int) (*model.StudentListFromClass, error)
	// FEED
	CreatePost(classID int, newPost *model.PostCreate) (*model.Post, error)
	GetClassFeed(classID int) (*model.Feed, error)
	// HOMEWORK
	CreateHomework(newHw *model.HomeworkCreate) (*model.Homework, error)
	GetHomeworkByID(id int) (*model.HomeworkByID, error)
	GetHomeworksByClassID(classID int) (*model.HomeworkList, error)
	// SOLUTION
	GetSolutionByID(id int) (*model.SolutionByID, error)
	GetSolutionsByClassID(classID int) (*model.SolutionListFromClass, error)
	GetSolutionsByHwID(hwID int) (*model.SolutionListForHw, error)
	// CALENDAR
	SetOAUTH2Token() error
	SaveOAUTH2Token(authCode string) error
	CreateCalendar(teacherID int) (*model.CreateCalendarResponse, error)
	CreateCalendarEvent(req *model.CreateCalendarEvent, teacherID int, classID int) error
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
