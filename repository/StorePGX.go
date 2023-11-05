package repository

import (
	"main/domain/model"

	e "main/domain/errors"

	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type StoreInterface interface {
	AddTeacher(in *model.TeacherSignUp) error
	GetTeacherProfile(id int) (*model.TeacherProfile, error)
	GetChatByID(id int) (*model.Chat, error)
	GetChatsByTeacherID(idTeacher int) (*model.ChatsPreview, error)
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) StoreInterface {
	return &Store{
		db: db,
	}
}

func (us *Store) AddTeacher(in *model.TeacherSignUp) error {
	_, err := us.db.Exec(
		`INSERT INTO teachers (login, name, password) VALUES ($1, $2, $3);`,
		in.Login, in.Name, in.Password,
	)
	if err != nil {
		return e.StacktraceError(err)
	}

	return nil
}

func (us *Store) GetTeacherProfile(id int) (*model.TeacherProfile, error) {
	teacher := &model.TeacherProfile{}
	rows, err := us.db.Query(
		`SELECT name FROM teachers WHERE id = $1`,
		id,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&teacher.Name)
		if err != nil {
			return nil, e.StacktraceError(err)
		}
	}

	return teacher, nil
}

func (us *Store) GetChatByID(id int) (*model.Chat, error) {
	messages := []*model.Message{}
	tmp := 1
	row := us.db.QueryRow(
		`SELECT 1 FROM chats WHERE id = $1`,
		id,
	)
	err := row.Scan(&tmp)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	rows, err := us.db.Query(
		`SELECT id, text, isAuthorTeacher, attaches, time, isRead FROM messages
		 WHERE chatID = $1`,
		id,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	for rows.Next() {
		dat := model.Message{}
		err := rows.Scan(
			&dat.ID, &dat.Text, &dat.IsAuthorTeacher,
			&dat.Attaches, &dat.Time, &dat.IsRead,
		)
		if err != nil {
			return nil, e.StacktraceError(err)
		}

		messages = append(messages, &dat)
	}
	return &model.Chat{Messages: messages}, nil
}

func (us *Store) GetChatsByTeacherID(idTeacher int) (*model.ChatsPreview, error) {
	chats := []*model.ChatPreview{}

	rows, err := us.db.Query(
		`SELECT id FROM chats WHERE teacherID = $1`,
		idTeacher,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var tmpID int
		err := rows.Scan(&tmpID)
		if err != nil {
			return nil, e.StacktraceError(err)
		}

		tmpChat := model.ChatPreview{
			ChatID: tmpID,
			Name:   "mockName",
			Img:    "mockImg",
		}

		row := us.db.QueryRow(
			`SELECT text, time, isRead FROM messages
			 WHERE chatID = $1
			 ORDER BY id DESC
			 LIMIT 1`,
			tmpID,
		)

		err = row.Scan(&tmpChat.LastMessageText, &tmpChat.LastMessageDate, &tmpChat.IsRead)
		if err != nil {
			return nil, e.StacktraceError(err)
		}
		chats = append(chats, &tmpChat)
	}

	return &model.ChatsPreview{Chats: chats}, nil
}
