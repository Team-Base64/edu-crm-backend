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
	GetChatsByTeacherID(teacherID int) (*model.ChatPreviewList, error)
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
