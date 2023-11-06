package repository

import (
	"main/domain/model"

	e "main/domain/errors"

	"database/sql"

	"github.com/jackc/pgtype"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type StoreInterface interface {
	CheckChatExistence(id int) error
	AddTeacher(in *model.TeacherSignUp) error
	GetTeacherProfile(id int) (*model.TeacherProfile, error)
	GetChatByID(id int) (*model.Chat, error)
	GetChatsByTeacherID(idTeacher int) (*model.ChatPreviewList, error)
	GetClassesByID(teacherID int) (*model.ClassesInfo, error)
	GetClassByID(id int) (*model.ClassInfo, error)
	AddClass(teacherID int, inviteToken string, newClass *model.ClassCreate) (int, error)
	GetStudentsFromClass(classID int) (*model.StudentsFromClass, error)
	GetClassFeed(classID int) (*model.Feed, error)
	GetHomeworksByClassID(classID int) (*model.HomeworksFromClass, error)
	GetHomeworkByID(id int) (*model.HomeworkByID, error)
	GetSolutionsByClassID(classID int) (*model.SolutionsFromClass, error)
	GetSolutionsByHwID(hwID int) (*model.SolutionsForHw, error)
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
	row := s.db.QueryRow(
		`SELECT 1 FROM chats WHERE id = $1;`,
		id,
	)
	if err := row.Scan(&tmp); err != nil {
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
	row := s.db.QueryRow(
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
		`SELECT id, text, isAuthorTeacher, attaches, time, isRead FROM messages
		 WHERE chatID = $1;`,
		id,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	var messages []*model.Message
	for rows.Next() {
		var tmpMsg model.Message
		var tmpAttaches pgtype.TextArray

		if err := rows.Scan(
			&tmpMsg.ID, &tmpMsg.Text,
			&tmpMsg.IsAuthorTeacher, &tmpAttaches,
			&tmpMsg.Time, &tmpMsg.IsRead,
		); err != nil {
			return nil, e.StacktraceError(err)
		}

		tmpMsg.Attaches = make([]string, len(tmpAttaches.Elements))
		for idx, el := range tmpAttaches.Elements {
			tmpMsg.Attaches[idx] = el.String
		}

		messages = append(messages, &tmpMsg)
	}
	return &model.Chat{Messages: messages}, nil
}

func (s *Store) GetChatsByTeacherID(teacherID int) (*model.ChatPreviewList, error) {
	rows, err := s.db.Query(
		`SELECT id FROM chats WHERE teacherID = $1;`,
		teacherID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	var chats []*model.ChatPreview
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
			`SELECT text, time, isRead FROM messages
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

func (us *Store) GetClassesByID(teacherID int) (*model.ClassesInfo, error) {
	rows, err := us.db.Query(
		`SELECT id, title, description, inviteToken FROM classes WHERE teacherID = $1;`,
		teacherID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	classes := []*model.ClassInfo{}
	for rows.Next() {
		tmpClass := model.ClassInfo{}
		err := rows.Scan(&tmpClass.ID, &tmpClass.Title, &tmpClass.Description, &tmpClass.InviteToken)
		if err != nil {
			return nil, e.StacktraceError(err)
		}
		classes = append(classes, &tmpClass)
	}

	return &model.ClassesInfo{Classes: classes}, nil
}

func (us *Store) GetClassByID(id int) (*model.ClassInfo, error) {
	row := us.db.QueryRow(
		`SELECT title, description, inviteToken FROM classes WHERE id = $1;`,
		id,
	)
	class := model.ClassInfo{}
	err := row.Scan(&class.Title, &class.Description, &class.InviteToken)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	class.ID = id
	return &class, nil
}

func (us *Store) AddClass(teacherID int, inviteToken string, newClass *model.ClassCreate) (int, error) {
	var id int

	row := us.db.QueryRow(
		`INSERT INTO classes (teacherID, title, description, inviteToken)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id;`,
		teacherID, newClass.Title, newClass.Description, inviteToken,
	)
	err := row.Scan(&id)
	if err != nil {
		return 0, e.StacktraceError(err)
	}

	return int(id), nil
}

func (us *Store) GetStudentsFromClass(classID int) (*model.StudentsFromClass, error) {
	var tmp int
	row := us.db.QueryRow(
		`SELECT 1 FROM classes WHERE id = $1`,
		classID,
	)
	err := row.Scan(&tmp)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	rows, err := us.db.Query(
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
		tmpStudent := model.Student{}
		err := rows.Scan(&tmpStudent.ID, &tmpStudent.Name, &tmpStudent.SocialType)
		if err != nil {
			return nil, e.StacktraceError(err)
		}
		students = append(students, &tmpStudent)
	}

	return &model.StudentsFromClass{Students: students}, nil
}

func (us *Store) GetClassFeed(classID int) (*model.Feed, error) {
	var tmp int
	row := us.db.QueryRow(
		`SELECT 1 FROM classes WHERE id = $1`,
		classID,
	)
	err := row.Scan(&tmp)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	rows, err := us.db.Query(
		`SELECT id, text, attaches, time FROM posts WHERE classID = $1;`,
		classID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	posts := []*model.Post{}
	for rows.Next() {
		tmpPost := model.Post{}
		tmpAttaches := pgtype.TextArray{}
		err := rows.Scan(&tmpPost.ID, &tmpPost.Text, &tmpAttaches, &tmpPost.Time)
		if err != nil {
			return nil, e.StacktraceError(err)
		}

		tmpPost.Attaches = make([]string, len(tmpAttaches.Elements))
		for idx, el := range tmpAttaches.Elements {
			tmpPost.Attaches[idx] = el.String
		}

		posts = append(posts, &tmpPost)
	}

	return &model.Feed{Posts: posts}, nil
}

func (us *Store) GetHomeworksByClassID(classID int) (*model.HomeworksFromClass, error) {
	var tmp int
	row := us.db.QueryRow(
		`SELECT 1 FROM classes WHERE id = $1`,
		classID,
	)
	err := row.Scan(&tmp)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	rows, err := us.db.Query(
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
		tmpHw := model.HomeworkFromClass{}
		err := rows.Scan(
			&tmpHw.ID, &tmpHw.Title, &tmpHw.Description,
			&tmpHw.CreateTime, &tmpHw.DeadlineTime, &tmpHw.File,
		)
		if err != nil {
			return nil, e.StacktraceError(err)
		}
		hws = append(hws, &tmpHw)
	}

	return &model.HomeworksFromClass{Homeworks: hws}, nil
}

func (us *Store) GetHomeworkByID(id int) (*model.HomeworkByID, error) {
	row := us.db.QueryRow(
		`SELECT classID, title, description, createTime, deadlineTime, file
		 FROM homeworks
		 WHERE id = $1;`,
		id,
	)
	hw := model.HomeworkByID{}
	err := row.Scan(
		&hw.ClassID, &hw.Title, &hw.Description,
		&hw.CreateTime, &hw.DeadlineTime, &hw.File,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	return &hw, nil
}

func (us *Store) GetSolutionsByClassID(classID int) (*model.SolutionsFromClass, error) {
	var tmp int
	row := us.db.QueryRow(
		`SELECT 1 FROM classes WHERE id = $1`,
		classID,
	)
	err := row.Scan(&tmp)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	rows, err := us.db.Query(
		`SELECT s.id, s.hwID, s.studentID, s.text, s.time, s.file
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
		tmpSol := model.SolutionFromClass{}
		err := rows.Scan(
			&tmpSol.ID, &tmpSol.HwID, &tmpSol.StudentID,
			&tmpSol.Text, &tmpSol.Time, &tmpSol.File,
		)
		if err != nil {
			return nil, e.StacktraceError(err)
		}
		sols = append(sols, &tmpSol)
	}

	return &model.SolutionsFromClass{Solutions: sols}, nil
}

func (us *Store) GetSolutionsByHwID(hwID int) (*model.SolutionsForHw, error) {
	var tmp int
	row := us.db.QueryRow(
		`SELECT 1 FROM homeworks WHERE id = $1`,
		hwID,
	)
	err := row.Scan(&tmp)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	rows, err := us.db.Query(
		`SELECT id, studentID, text, time, file FROM solutions WHERE hwID = $1;`,
		hwID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	sols := []*model.SolutionForHw{}
	for rows.Next() {
		tmpSol := model.SolutionForHw{}
		err := rows.Scan(
			&tmpSol.ID, &tmpSol.StudentID, &tmpSol.Text, &tmpSol.Time, &tmpSol.File,
		)
		if err != nil {
			return nil, e.StacktraceError(err)
		}
		sols = append(sols, &tmpSol)
	}

	return &model.SolutionsForHw{Solutions: sols}, nil
}

func (us *Store) GetSolutionByID(id int) (*model.SolutionByID, error) {
	row := us.db.QueryRow(
		`SELECT hwID, studentID, text, time, file FROM solutions WHERE id = $1;`,
		id,
	)
	sol := model.SolutionByID{}
	err := row.Scan(&sol.HwID, &sol.StudentID, &sol.Text, &sol.Time, &sol.File)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	return &sol, nil
}
