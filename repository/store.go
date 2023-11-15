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
	// CHAT
	CheckChatExistence(id int) error
	GetChatByID(id int) (*model.Chat, error)
	GetChatsByTeacherID(idTeacher int) (*model.ChatPreviewList, error)
	// CLASS
	CheckClassExistence(id int) error
	AddClass(teacherID int, inviteToken string, newClass *model.ClassCreate) (int, error)
	GetClassByID(id int) (*model.ClassInfo, error)
	GetClassesByID(teacherID int) (*model.ClassInfoList, error)
	// STUDENT
	GetStudentByID(id int) (*model.StudentByID, error)
	GetStudentsFromClass(classID int) (*model.StudentListFromClass, error)
	// FEED
	AddPost(classID int, createTime time.Time, newPost *model.PostCreate) (int, error)
	DeletePost(id int) error
	GetClassFeed(classID int) (*model.Feed, error)
	// HOMEWORK
	CheckHomeworkExistence(id int) error
	AddHomework(teacherID int, createTime time.Time, newHw *model.HomeworkCreate) (int, error)
	DeleteHomework(id int) error
	GetHomeworkByID(id int) (*model.HomeworkByID, error)
	GetHomeworksByClassID(classID int) ([]model.Homework, error)
	// TASK
	AddTask(teacherID int, newTask *model.TaskCreate) (int, error)
	GetTaskByID(id int) (*model.TaskByID, error)
	GetTasksByTeacherID(teacherID int) ([]*model.Task, error)
	GetTasksByHomeworkID(homeworkID int) ([]*model.Task, error)
	GetTasksIDByHomeworkID(homeworkID int) ([]int, error)
	AttachTaskToHomework(hwID int, taskID int, taskRank int) error
	// SOLUTION
	GetSolutionByID(id int) (*model.SolutionByID, error)
	GetSolutionsByClassID(classID int) (*model.SolutionListFromClass, error)
	GetSolutionsByHomeworkID(homeworkID int) (*model.SolutionListForHw, error)
	// CALENDAR
	GetTokenDB(id int) (string, error)
	CreateCalendarDB(teacherID int, googleID string) (int, error)
	GetCalendarGoogleID(teacherID int) (string, error)
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) StoreInterface {
	return &Store{
		db: db,
	}
}
