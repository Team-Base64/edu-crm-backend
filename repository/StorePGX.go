package repository

import (
	"main/domain/model"
	"time"

	e "main/domain/errors"

	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/lib/pq"
)

type StoreInterface interface {
	CheckChatExistence(id int) error
	CheckClassExistence(id int) error
	CheckHomeworkExistence(id int) error
	AddTeacher(in *model.TeacherSignUp) error
	GetTeacherProfile(id int) (*model.TeacherProfile, error)
	GetChatByID(id int) (*model.Chat, error)
	GetChatsByTeacherID(idTeacher int) (*model.ChatPreviewList, error)
	GetClassesByID(teacherID int) (*model.ClassInfoList, error)
	GetClassByID(id int) (*model.ClassInfo, error)
	AddClass(teacherID int, inviteToken string, newClass *model.ClassCreate) (int, error)
	GetStudentByID(id int) (*model.StudentByID, error)
	GetStudentsFromClass(classID int) (*model.StudentListFromClass, error)
	GetClassFeed(classID int) (*model.Feed, error)
	AddPost(classID int, createTime time.Time, newPost *model.PostCreate) (int, error)
	DeletePost(id int) error
	GetHomeworksByClassID(classID int) (*model.HomeworkList, error)
	GetHomeworkByID(id int) (*model.HomeworkByID, error)
	AddHomework(createTime time.Time, newHw *model.HomeworkCreate) (int, error)
	DeleteHomework(id int) error
	GetTaskByID(id int) (*model.TaskByID, error)
	AddTask(newTask *model.TaskCreate) (int, error)
	AttachTaskToHomework(hwID int, taskID int) error
	GetSolutionsByClassID(classID int) (*model.SolutionListFromClass, error)
	GetSolutionsByHwID(hwID int) (*model.SolutionListForHw, error)
	GetSolutionByID(id int) (*model.SolutionByID, error)
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) StoreInterface {
	return &Store{
		db: db,
	}
}

func (s *Store) CheckChatExistence(id int) error {
	var tmp int
	if err := s.db.QueryRow(
		`SELECT 1 FROM chats WHERE id = $1;`,
		id,
	).Scan(&tmp); err != nil {
		return e.StacktraceError(err)
	}
	return nil
}

func (s *Store) CheckClassExistence(id int) error {
	var tmp int
	if err := s.db.QueryRow(
		`SELECT 1 FROM classes WHERE id = $1;`,
		id,
	).Scan(&tmp); err != nil {
		return e.StacktraceError(err)
	}
	return nil
}

func (s *Store) CheckHomeworkExistence(id int) error {
	var tmp int
	if err := s.db.QueryRow(
		`SELECT 1 FROM homeworks WHERE id = $1;`,
		id,
	).Scan(&tmp); err != nil {
		return e.StacktraceError(err)
	}
	return nil
}

func (s *Store) AddTeacher(in *model.TeacherSignUp) error {
	_, err := s.db.Exec(
		`INSERT INTO teachers (login, name, password) VALUES ($1, $2, $3);`,
		in.Login, in.Name, in.Password,
	)
	if err != nil {
		return e.StacktraceError(err)
	}

	return nil
}

func (s *Store) GetTeacherProfile(id int) (*model.TeacherProfile, error) {
	var teacher model.TeacherProfile
	if err := s.db.QueryRow(
		`SELECT name FROM teachers WHERE id = $1;`,
		id,
	).Scan(&teacher.Name); err != nil {
		return nil, e.StacktraceError(err)
	}

	return &teacher, nil
}

func (s *Store) GetChatByID(id int) (*model.Chat, error) {
	rows, err := s.db.Query(
		`SELECT id, text, isAuthorTeacher, attaches, createTime, isRead FROM messages
		 WHERE chatID = $1;`,
		id,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	messages := []*model.Message{}
	for rows.Next() {
		var tmpMsg model.Message

		if err := rows.Scan(
			&tmpMsg.ID, &tmpMsg.Text,
			&tmpMsg.IsAuthorTeacher, (*pq.StringArray)(&tmpMsg.Attaches),
			&tmpMsg.CreateTime, &tmpMsg.IsRead,
		); err != nil {
			return nil, e.StacktraceError(err)
		}

		messages = append(messages, &tmpMsg)
	}
	return &model.Chat{Messages: messages}, nil
}

func (s *Store) GetChatsByTeacherID(teacherID int) (*model.ChatPreviewList, error) {
	rows, err := s.db.Query(
		`SELECT m1.chatID, s.name, s.socialType, m1.text, m1.createTime, m1.isRead
		 FROM messages m1
		 LEFT JOIN messages m2
		 ON m1.chatId = m2.chatId AND m1.createTime < m2.createTime
		 JOIN chats c ON m1.chatID = c.id
		 JOIN students s ON c.studentID = s.id
		 WHERE m2.chatID IS NULL AND c.teacherID = $1 ORDER BY m1.createTime DESC;`,
		teacherID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	chats := []*model.ChatPreview{}
	for rows.Next() {
		tmpChat := model.ChatPreview{
			Img: "mockImg",
		}

		if err = rows.Scan(
			&tmpChat.ChatID,
			&tmpChat.Name,
			&tmpChat.SocialType,
			&tmpChat.LastMessageText,
			&tmpChat.LastMessageDate,
			&tmpChat.IsRead,
		); err != nil {
			return nil, e.StacktraceError(err)
		}
		chats = append(chats, &tmpChat)
	}

	return &model.ChatPreviewList{Chats: chats}, nil
}

func (s *Store) GetClassesByID(teacherID int) (*model.ClassInfoList, error) {
	rows, err := s.db.Query(
		`SELECT id, title, description, inviteToken FROM classes WHERE teacherID = $1;`,
		teacherID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	classes := []*model.ClassInfo{}
	for rows.Next() {
		var tmpClass model.ClassInfo

		if err := rows.Scan(
			&tmpClass.ID, &tmpClass.Title,
			&tmpClass.Description, &tmpClass.InviteToken,
		); err != nil {
			return nil, e.StacktraceError(err)
		}
		classes = append(classes, &tmpClass)
	}

	return &model.ClassInfoList{Classes: classes}, nil
}

func (s *Store) GetClassByID(id int) (*model.ClassInfo, error) {
	var class model.ClassInfo
	if err := s.db.QueryRow(
		`SELECT title, description, inviteToken FROM classes WHERE id = $1;`,
		id,
	).Scan(
		&class.Title,
		&class.Description,
		&class.InviteToken,
	); err != nil {
		return nil, e.StacktraceError(err)
	}

	class.ID = id
	return &class, nil
}

func (s *Store) AddClass(teacherID int, inviteToken string, newClass *model.ClassCreate) (int, error) {
	var id int
	if err := s.db.QueryRow(
		`INSERT INTO classes (teacherID, title, description, inviteToken)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id;`,
		teacherID, newClass.Title, newClass.Description, inviteToken,
	).Scan(&id); err != nil {
		return 0, e.StacktraceError(err)
	}

	return int(id), nil
}

func (s *Store) GetStudentByID(id int) (*model.StudentByID, error) {
	var stud model.StudentByID
	if err := s.db.QueryRow(
		`SELECT name, socialType FROM students WHERE id = $1;`,
		id,
	).Scan(
		&stud.Name, &stud.SocialType,
	); err != nil {
		return nil, e.StacktraceError(err)
	}

	return &stud, nil
}

func (s *Store) GetStudentsFromClass(classID int) (*model.StudentListFromClass, error) {
	rows, err := s.db.Query(
		`SELECT s.id, s.name, s.socialType FROM students s
		 JOIN classes_students cs ON s.id = cs.studentID
		 WHERE cs.classID = $1;`,
		classID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	students := []*model.StudentFromClass{}
	for rows.Next() {
		var tmpStudent model.StudentFromClass

		if err := rows.Scan(
			&tmpStudent.ID,
			&tmpStudent.Name,
			&tmpStudent.SocialType,
		); err != nil {
			return nil, e.StacktraceError(err)
		}
		students = append(students, &tmpStudent)
	}

	return &model.StudentListFromClass{Students: students}, nil
}

func (s *Store) GetClassFeed(classID int) (*model.Feed, error) {
	rows, err := s.db.Query(
		`SELECT id, text, attaches, createTime FROM posts WHERE classID = $1;`,
		classID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	posts := []*model.Post{}
	for rows.Next() {
		var tmpPost model.Post

		if err := rows.Scan(
			&tmpPost.ID, &tmpPost.Text,
			(*pq.StringArray)(&tmpPost.Attaches), &tmpPost.CreateTime,
		); err != nil {
			return nil, e.StacktraceError(err)
		}

		posts = append(posts, &tmpPost)
	}

	return &model.Feed{Posts: posts}, nil
}

func (s *Store) AddPost(classID int, createTime time.Time, newPost *model.PostCreate) (int, error) {
	var id int
	if err := s.db.QueryRow(
		`INSERT INTO posts (classID, text, attaches, createTime)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id;`,
		classID, newPost.Text, (*pq.StringArray)(&newPost.Attaches), createTime,
	).Scan(&id); err != nil {
		return 0, e.StacktraceError(err)
	}

	return int(id), nil
}

func (s *Store) DeletePost(id int) error {
	_, err := s.db.Exec(
		`DELETE FROM posts WHERE id = $1;`,
		id,
	)

	if err != nil {
		return e.StacktraceError(err)
	}

	return nil
}

func (s *Store) GetHomeworksByClassID(classID int) (*model.HomeworkList, error) {
	rows, err := s.db.Query(
		`SELECT id, title, description, createTime, deadlineTime
		 FROM homeworks
		 WHERE classID = $1;`,
		classID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	hws := []*model.Homework{}
	for rows.Next() {
		var tmpHw model.Homework

		if err := rows.Scan(
			&tmpHw.ID, &tmpHw.Title, &tmpHw.Description,
			&tmpHw.CreateTime, &tmpHw.DeadlineTime,
		); err != nil {
			return nil, e.StacktraceError(err)
		}

		hws = append(hws, &tmpHw)
	}

	return &model.HomeworkList{Homeworks: hws}, nil
}

func (s *Store) GetHomeworkByID(id int) (*model.HomeworkByID, error) {
	var hw model.HomeworkByID
	if err := s.db.QueryRow(
		`SELECT classID, title, description, createTime, deadlineTime, file
		 FROM homeworks
		 WHERE id = $1;`,
		id,
	).Scan(
		&hw.ClassID, &hw.Title, &hw.Description,
		&hw.CreateTime, &hw.DeadlineTime, &hw.File,
	); err != nil {
		return nil, e.StacktraceError(err)
	}

	return &hw, nil
}

func (s *Store) AddHomework(createTime time.Time, newHw *model.HomeworkCreate) (int, error) {
	var id int
	if err := s.db.QueryRow(
		`INSERT INTO homeworks (classID, title, description, createTime, deadlineTime)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING id;`,
		newHw.ClassID, newHw.Title, newHw.Description, createTime, newHw.DeadlineTime,
	).Scan(&id); err != nil {
		return 0, e.StacktraceError(err)
	}

	for _, task := range newHw.Tasks {
		taskID := task.ID
		if taskID < 0 {
			newTaskID, err := s.AddTask(&model.TaskCreate{
				Description: task.Description,
				Attach:      task.Attach,
			})
			if err != nil {
				return 0, e.StacktraceError(err)
			}
			taskID = newTaskID
		}
		s.AttachTaskToHomework(id, taskID)
	}

	return int(id), nil
}

func (s *Store) DeleteHomework(id int) error {
	_, err := s.db.Exec(
		`DELETE FROM homeworks WHERE id = $1;`,
		id,
	)

	if err != nil {
		return e.StacktraceError(err)
	}

	return nil
}

func (s *Store) GetTaskByID(id int) (*model.TaskByID, error) {
	var task model.TaskByID
	if err := s.db.QueryRow(
		`SELECT description, attach FROM solutions WHERE id = $1;`,
		id,
	).Scan(
		&task.Description, &task.Attach,
	); err != nil {
		return nil, e.StacktraceError(err)
	}

	return &task, nil
}

func (s *Store) AddTask(newTask *model.TaskCreate) (int, error) {
	var id int
	if err := s.db.QueryRow(
		`INSERT INTO tasks (description, attach)
		 VALUES ($1, $2)
		 RETURNING id;`,
		newTask.Description, newTask.Attach,
	).Scan(&id); err != nil {
		return 0, e.StacktraceError(err)
	}

	return int(id), nil
}

func (s *Store) AttachTaskToHomework(hwID int, taskID int) error {
	_, err := s.db.Exec(
		`INSERT INTO homeworks_tasks (homeworkID, taskID) VALUES ($1, $2)`,
		hwID, taskID,
	)
	if err != nil {
		return e.StacktraceError(err)
	}

	return nil
}

func (s *Store) GetSolutionsByClassID(classID int) (*model.SolutionListFromClass, error) {
	rows, err := s.db.Query(
		`SELECT s.id, s.hwID, s.studentID, s.text, s.createTime, s.file
		 FROM solutions s
		 JOIN homeworks h ON s.hwID = h.id
		 WHERE h.classID = $1;`,
		classID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	sols := []*model.SolutionFromClass{}
	for rows.Next() {
		var tmpSol model.SolutionFromClass

		if err := rows.Scan(
			&tmpSol.ID, &tmpSol.HwID, &tmpSol.StudentID,
			&tmpSol.Text, &tmpSol.CreateTime, &tmpSol.File,
		); err != nil {
			return nil, e.StacktraceError(err)
		}

		sols = append(sols, &tmpSol)
	}

	return &model.SolutionListFromClass{Solutions: sols}, nil
}

func (s *Store) GetSolutionsByHwID(hwID int) (*model.SolutionListForHw, error) {
	rows, err := s.db.Query(
		`SELECT id, studentID, text, createTime, file FROM solutions WHERE hwID = $1;`,
		hwID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	sols := []*model.SolutionForHw{}
	for rows.Next() {
		var tmpSol model.SolutionForHw

		if err := rows.Scan(
			&tmpSol.ID, &tmpSol.StudentID,
			&tmpSol.Text, &tmpSol.CreateTime, &tmpSol.File,
		); err != nil {
			return nil, e.StacktraceError(err)
		}

		sols = append(sols, &tmpSol)
	}

	return &model.SolutionListForHw{Solutions: sols}, nil
}

func (s *Store) GetSolutionByID(id int) (*model.SolutionByID, error) {
	var sol model.SolutionByID
	if err := s.db.QueryRow(
		`SELECT hwID, studentID, text, createTime, file FROM solutions WHERE id = $1;`,
		id,
	).Scan(
		&sol.HwID, &sol.StudentID,
		&sol.Text, &sol.CreateTime, &sol.File,
	); err != nil {
		return nil, e.StacktraceError(err)
	}

	return &sol, nil
}
