package repository

import (
	"time"

	m "main/domain/model"
)

type DataStoreInterface interface {
	// TEACHER
	AddTeacher(in *m.TeacherSignUp) error
	GetTeacherProfile(id int) (*m.TeacherProfile, error)
	GetTeacherProfileByLoginDB(login string) (*m.TeacherDB, error)
	// CHAT
	CheckChatExistence(id int) error
	GetChatByID(id int) (*m.Chat, error)
	GetChatsByTeacherID(idTeacher int) ([]m.ChatPreview, error)
	GetChatIDBySolutionID(solutionID int) (int, error)
	ReadChatByID(id int, teacherID int) error
	// CLASS
	CheckClassExistence(id int) error
	AddClass(teacherID int, inviteToken string, newClass *m.ClassCreate) (int, error)
	GetClassByID(id int) (*m.ClassInfo, error)
	GetClassesByTeacherID(teacherID int) ([]m.ClassInfo, error)
	// STUDENT
	GetStudentByID(id int) (*m.StudentByID, error)
	GetStudentsFromClass(classID int) ([]m.StudentFromClass, error)
	// FEED
	AddPost(classID int, createTime time.Time, newPost *m.PostCreate) (int, error)
	DeletePost(id int) error
	GetClassPosts(classID int) ([]m.Post, error)
	// HOMEWORK
	CheckHomeworkExistence(id int) error
	AddHomework(teacherID int, createTime time.Time, newHw *m.HomeworkCreate) (int, error)
	DeleteHomework(id int) error
	GetHomeworkByID(id int) (*m.HomeworkByID, error)
	GetHomeworksByClassID(classID int) ([]m.Homework, error)
	// TASK
	AddTask(teacherID int, newTask *m.TaskCreate) (int, error)
	GetTaskByID(id int) (*m.TaskByID, error)
	GetTasksByTeacherID(teacherID int) ([]m.Task, error)
	GetTasksByHomeworkID(homeworkID int) ([]m.Task, error)
	GetTasksIDByHomeworkID(homeworkID int) ([]int, error)
	AttachTaskToHomework(hwID int, taskID int, taskRank int) error
	// SOLUTION
	GetSolutionByID(id int) (*m.SolutionByID, error)
	GetSolutionsByClassID(classID int) ([]m.SolutionFromClass, error)
	GetSolutionsByHomeworkID(homeworkID int) ([]m.SolutionForHw, error)
	GetInfoForEvaluationMsgBySolutionID(solutionID int) (*m.SolutionInfoForEvaluationMsg, error)
	AddEvaluationForSolution(solutionID int, isApproved bool, evaluation string) error
	// CALENDAR
	GetCalendar(teacherID int) (*m.CalendarParams, error)
	// SESSIONS
	CreateSession(teacherLogin string) (*m.Session, error)
	CheckSession(in string) (string, error)
	DeleteSession(in string) error
}

type FileStoreInterface interface {
	UploadFile(file *m.Attach) (string, error)
}
