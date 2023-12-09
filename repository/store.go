package repository

import (
	"main/domain/model"
	"time"

	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type StoreInterface interface {
	// TEACHER
	AddTeacher(in *model.TeacherSignUp) error
	GetTeacherProfile(id int) (*model.TeacherProfile, error)
	GetTeacherProfileByLoginDB(login string) (*model.TeacherDB, error)
	// CHAT
	CheckChatExistence(id int) error
	GetChatByID(id int) (*model.Chat, error)
	GetChatsByTeacherID(idTeacher int) ([]model.ChatPreview, error)
	GetChatIDBySolutionID(solutionID int) (int, error)
	ReadChatByID(id int, teacherID int) error
	// CLASS
	CheckClassExistence(id int) error
	AddClass(teacherID int, inviteToken string, newClass *model.ClassCreate) (int, error)
	GetClassByID(id int) (*model.ClassInfo, error)
	GetClassesByTeacherID(teacherID int) ([]model.ClassInfo, error)
	// STUDENT
	GetStudentByID(id int) (*model.StudentByID, error)
	GetStudentsFromClass(classID int) ([]model.StudentFromClass, error)
	// FEED
	AddPost(classID int, createTime time.Time, newPost *model.PostCreate) (int, error)
	DeletePost(id int) error
	GetClassPosts(classID int) ([]model.Post, error)
	// HOMEWORK
	CheckHomeworkExistence(id int) error
	AddHomework(teacherID int, createTime time.Time, newHw *model.HomeworkCreate) (int, error)
	DeleteHomework(id int) error
	GetHomeworkByID(id int) (*model.HomeworkByID, error)
	GetHomeworksByClassID(classID int) ([]model.Homework, error)
	// TASK
	AddTask(teacherID int, newTask *model.TaskCreate) (int, error)
	GetTaskByID(id int) (*model.TaskByID, error)
	GetTasksByTeacherID(teacherID int) ([]model.Task, error)
	GetTasksByHomeworkID(homeworkID int) ([]model.Task, error)
	GetTasksIDByHomeworkID(homeworkID int) ([]int, error)
	AttachTaskToHomework(hwID int, taskID int, taskRank int) error
	// SOLUTION
	GetSolutionByID(id int) (*model.SolutionByID, error)
	GetSolutionsByClassID(classID int) ([]model.SolutionFromClass, error)
	GetSolutionsByHomeworkID(homeworkID int) ([]model.SolutionForHw, error)
	GetInfoForEvaluationMsgBySolutionID(solutionID int) (*model.SolutionInfoForEvaluationMsg, error)
	AddEvaluationForSolution(solutionID int, isApproved bool, evaluation string) error
	// CALENDAR
	GetTokenDB(id int) (string, error)
	CreateCalendarDB(teacherID int, googleID string) (int, error)
	GetCalendarDB(teacherID int) (*model.CalendarParams, error)
	// SESSIONS
	CreateSession(teacherLogin string) (*model.Session, error)
	CheckSession(in string) (string, error)
	DeleteSession(in string) error
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) StoreInterface {
	return &Store{
		db: db,
	}
}
