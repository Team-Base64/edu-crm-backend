package repository

import (
	"context"
	"main/domain/model"
	"time"

	e "main/domain/errors"

	"github.com/jackc/pgx/v5"
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
	GetStudentsFromClass(classID int) (*model.StudentListFromClass, error)
	GetClassFeed(classID int) (*model.Feed, error)
	AddPost(classID int, createTime time.Time, newPost *model.PostCreate) (int, error)
	DeletePost(id int) error
	GetHomeworksByClassID(classID int) (*model.HomeworkListFromClass, error)
	GetHomeworkByID(id int) (*model.HomeworkByID, error)
	AddHomework(createTime time.Time, newHw *model.HomeworkCreate) (int, error)
	DeleteHomework(id int) error
	GetSolutionsByClassID(classID int) (*model.SolutionListFromClass, error)
	GetSolutionsByHwID(hwID int) (*model.SolutionListForHw, error)
	GetSolutionByID(id int) (*model.SolutionByID, error)
}

type Store struct {
	db *pgx.Conn
}

func NewStore(db *pgx.Conn) StoreInterface {
	return &Store{
		db: db,
	}
}

func (s *Store) CheckChatExistence(id int) error {
	var tmp int
	row := s.db.QueryRow(
		context.Background(),
		`SELECT 1 FROM chats WHERE id = $1;`,
		id,
	)
	if err := row.Scan(&tmp); err != nil {
		return e.StacktraceError(err)
	}
	return nil
}

func (s *Store) CheckClassExistence(id int) error {
	var tmp int
	row := s.db.QueryRow(
		context.Background(),
		`SELECT 1 FROM classes WHERE id = $1;`,
		id,
	)
	if err := row.Scan(&tmp); err != nil {
		return e.StacktraceError(err)
	}
	return nil
}

func (s *Store) CheckHomeworkExistence(id int) error {
	var tmp int
	row := s.db.QueryRow(
		context.Background(),
		`SELECT 1 FROM homeworks WHERE id = $1;`,
		id,
	)
	if err := row.Scan(&tmp); err != nil {
		return e.StacktraceError(err)
	}
	return nil
}

func (s *Store) AddTeacher(in *model.TeacherSignUp) error {
	_, err := s.db.Exec(
		context.Background(),
		`INSERT INTO teachers (login, name, password) VALUES ($1, $2, $3);`,
		in.Login, in.Name, in.Password,
	)
	if err != nil {
		return e.StacktraceError(err)
	}

	return nil
}

func (s *Store) GetTeacherProfile(id int) (*model.TeacherProfile, error) {
	row := s.db.QueryRow(
		context.Background(),
		`SELECT name FROM teachers WHERE id = $1;`,
		id,
	)

	var teacher model.TeacherProfile
	if err := row.Scan(&teacher.Name); err != nil {
		return nil, e.StacktraceError(err)
	}

	return &teacher, nil
}

func (s *Store) GetChatByID(id int) (*model.Chat, error) {
	rows, err := s.db.Query(
		context.Background(),
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
		// var tmpAttaches pgtype.TextArray

		if err := rows.Scan(
			&tmpMsg.ID, &tmpMsg.Text,
			&tmpMsg.IsAuthorTeacher, &tmpMsg.Attaches,
			&tmpMsg.CreateTime, &tmpMsg.IsRead,
		); err != nil {
			return nil, e.StacktraceError(err)
		}

		// tmpMsg.Attaches = make([]string, len(tmpAttaches.Elements))
		// for idx, el := range tmpAttaches.Elements {
		// 	tmpMsg.Attaches[idx] = el.String
		// }

		messages = append(messages, &tmpMsg)
	}
	return &model.Chat{Messages: messages}, nil
}

func (s *Store) GetChatsByTeacherID(teacherID int) (*model.ChatPreviewList, error) {
	rows, err := s.db.Query(
		context.Background(),
		`SELECT id FROM chats WHERE teacherID = $1;`,
		teacherID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	chats := []*model.ChatPreview{}
	for rows.Next() {
		var tmpID int
		if err := rows.Scan(&tmpID); err != nil {
			return nil, e.StacktraceError(err)
		}

		tmpChat := model.ChatPreview{
			ChatID: tmpID,
			Name:   "mockName",
			Img:    "mockImg",
		}

		row := s.db.QueryRow(
			context.Background(),
			`SELECT text, createTime, isRead FROM messages
			 WHERE chatID = $1
			 ORDER BY id DESC
			 LIMIT 1;`,
			tmpID,
		)

		if err = row.Scan(
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
		context.Background(),
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
	row := s.db.QueryRow(
		context.Background(),
		`SELECT title, description, inviteToken FROM classes WHERE id = $1;`,
		id,
	)
	var class model.ClassInfo

	if err := row.Scan(
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
	row := s.db.QueryRow(
		context.Background(),
		`INSERT INTO classes (teacherID, title, description, inviteToken)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id;`,
		teacherID, newClass.Title, newClass.Description, inviteToken,
	)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, e.StacktraceError(err)
	}

	return int(id), nil
}

func (s *Store) GetStudentsFromClass(classID int) (*model.StudentListFromClass, error) {
	rows, err := s.db.Query(
		context.Background(),
		`SELECT s.id, s.name, s.socialType FROM students s
		 JOIN classes_students cs ON s.id = cs.studentID
		 WHERE cs.classID = $1;`,
		classID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	students := []*model.Student{}
	for rows.Next() {
		var tmpStudent model.Student

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
		context.Background(),
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
		// var tmpAttaches pgtype.TextArray

		if err := rows.Scan(
			&tmpPost.ID, &tmpPost.Text,
			&tmpPost.Attaches, &tmpPost.CreateTime,
		); err != nil {
			return nil, e.StacktraceError(err)
		}

		// tmpPost.Attaches = make([]string, len(tmpAttaches.Elements))
		// for idx, el := range tmpAttaches.Elements {
		// 	tmpPost.Attaches[idx] = el.String
		// }

		posts = append(posts, &tmpPost)
	}

	return &model.Feed{Posts: posts}, nil
}

func (s *Store) AddPost(classID int, createTime time.Time, newPost *model.PostCreate) (int, error) {
	row := s.db.QueryRow(
		context.Background(),
		`INSERT INTO posts (classID, text, attaches, createTime)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id;`,
		classID, newPost.Text, newPost.Attaches, createTime,
	)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, e.StacktraceError(err)
	}

	return int(id), nil
}

func (s *Store) DeletePost(id int) error {
	_, err := s.db.Exec(
		context.Background(),
		`DELETE FROM posts WHERE id = $1;`,
		id,
	)

	if err != nil {
		return e.StacktraceError(err)
	}

	return nil
}

func (s *Store) GetHomeworksByClassID(classID int) (*model.HomeworkListFromClass, error) {
	rows, err := s.db.Query(
		context.Background(),
		`SELECT id, title, description, createTime, deadlineTime, file
		 FROM homeworks
		 WHERE classID = $1;`,
		classID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	hws := []*model.HomeworkFromClass{}
	for rows.Next() {
		var tmpHw model.HomeworkFromClass

		if err := rows.Scan(
			&tmpHw.ID, &tmpHw.Title, &tmpHw.Description,
			&tmpHw.CreateTime, &tmpHw.DeadlineTime, &tmpHw.File,
		); err != nil {
			return nil, e.StacktraceError(err)
		}

		hws = append(hws, &tmpHw)
	}

	return &model.HomeworkListFromClass{Homeworks: hws}, nil
}

func (s *Store) GetHomeworkByID(id int) (*model.HomeworkByID, error) {
	row := s.db.QueryRow(
		context.Background(),
		`SELECT classID, title, description, createTime, deadlineTime, file
		 FROM homeworks
		 WHERE id = $1;`,
		id,
	)

	var hw model.HomeworkByID
	if err := row.Scan(
		&hw.ClassID, &hw.Title, &hw.Description,
		&hw.CreateTime, &hw.DeadlineTime, &hw.File,
	); err != nil {
		return nil, e.StacktraceError(err)
	}

	return &hw, nil
}

func (s *Store) AddHomework(createTime time.Time, newHw *model.HomeworkCreate) (int, error) {
	row := s.db.QueryRow(
		context.Background(),
		`INSERT INTO homeworks (classID, title, description, createTime, deadlineTime, file)
		 VALUES ($1, $2, $3, $4, $5, $6)
		 RETURNING id;`,
		newHw.ClassID, newHw.Title, newHw.Description, createTime, newHw.DeadlineTime, newHw.File,
	)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, e.StacktraceError(err)
	}

	return int(id), nil
}

func (s *Store) DeleteHomework(id int) error {
	_, err := s.db.Exec(
		context.Background(),
		`DELETE FROM homeworks WHERE id = $1;`,
		id,
	)

	if err != nil {
		return e.StacktraceError(err)
	}

	return nil
}

func (s *Store) GetSolutionsByClassID(classID int) (*model.SolutionListFromClass, error) {
	rows, err := s.db.Query(
		context.Background(),
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
		context.Background(),
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
	row := s.db.QueryRow(
		context.Background(),
		`SELECT hwID, studentID, text, createTime, file FROM solutions WHERE id = $1;`,
		id,
	)
	var sol model.SolutionByID

	if err := row.Scan(
		&sol.HwID, &sol.StudentID,
		&sol.Text, &sol.CreateTime, &sol.File,
	); err != nil {
		return nil, e.StacktraceError(err)
	}

	return &sol, nil
}
