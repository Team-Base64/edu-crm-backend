package usecase

import (
	m "main/domain/model"
)

type UsecaseInterface interface {
	// TEACHER
	HashPass(plainPassword string) []byte
	CheckPass(passHash []byte, plainPassword string) bool
	SignUpTeacher(req *m.TeacherSignUp) error
	CreateTeacher(params *m.TeacherSignUp) error
	GetTeacherProfile(id int) (*m.TeacherProfile, error)
	GetTeacherProfileByLogin(login string) (*m.TeacherDB, error)
	// SESSION
	CreateSession(teacherLogin string) (*m.Session, error)
	CheckSession(in string) (string, error)
	DeleteSession(in string) error
	// CHAT
	GetChatByID(id int) (*m.Chat, error)
	GetChatsByTeacherID(id int) ([]m.ChatPreview, error)
	ReadChatByID(id int, teacherID int) error
	// CLASS
	CreateClass(teacherID int, newClass *m.ClassCreate) (*m.ClassInfo, error)
	GetClassByID(id int) (*m.ClassInfo, error)
	GetClassesByTeacherID(id int) ([]m.ClassInfo, error)
	// STUDENT
	GetStudentByID(id int) (*m.StudentByID, error)
	GetStudentsFromClass(classID int) ([]m.StudentFromClass, error)
	// FEED
	CreatePost(classID int, newPost *m.PostCreate) (*m.Post, error)
	GetClassPosts(classID int) ([]m.Post, error)
	// HOMEWORK
	CreateHomework(teacherID int, newHw *m.HomeworkCreate) (*m.Homework, error)
	GetHomeworkByID(id int) (*m.HomeworkByID, error)
	GetHomeworksByClassID(classID int) (*m.HomeworkList, error)
	// TASKS
	CreateTask(teacherID int, newTask *m.TaskCreate) (*m.TaskCreateResponse, error)
	GetTaskByID(id int) (*m.TaskByID, error)
	GetTasksByTeacherID(teacherID int) ([]m.Task, error)
	// SOLUTION
	GetSolutionByID(id int) (*m.SolutionByID, error)
	GetSolutionsByClassID(classID int) ([]m.SolutionFromClass, error)
	GetSolutionsByHomeworkID(homeworkID int) ([]m.SolutionForHw, error)
	EvaluateSolutionbyID(solutionID int, evaluation *m.SolutionEvaluation) error
	// CALENDAR
	CreateCalendar(teacherID int) error
	GetCalendarParams(teacherID int) (*m.CalendarParams, error)
	CreateCalendarEvent(req *m.CalendarEvent, teacherID int) error
	GetCalendarEvents(teacherID int) ([]m.CalendarEvent, error)
	DeleteCalendarEvent(teacherID int, eventID string) error
	UpdateCalendarEvent(req *m.CalendarEvent, teacherID int) error
	// ATTACH
	SaveAttach(file *m.Attach) (string, error)
}
