package repository

import (
	"context"
	"main/domain/model"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StoreInterface interface {
	AddTeacher(in *model.TeacherDB) error
	UpdateTeacher(in *model.TeacherDB) error
	GetTeacher(id int) (*model.TeacherDB, error)
	AddStudent(in *model.StudentDB) error
	UpdateStudent(in *model.StudentDB) error
	CreateChat(in *model.ChatDB) error
	AddMessage(in *model.CreateMessage) error
	//GetChatID(in *model.ChatDB) (int, error)
	GetChatFromDB(id int) (*model.Chat, error)
	GetChatsByID(idTeacher int) (*model.Chats, error)
}

type Store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) StoreInterface {
	return &Store{
		db: db,
	}
}

func (us *Store) AddTeacher(in *model.TeacherDB) error {
	_, err := us.db.Query(context.Background(), `INSERT INTO teachers (login, name, password) VALUES ($1, $2, $3);`, in.Login, in.Name, in.Password)
	if err != nil {
		return err
	}
	return nil
}

func (us *Store) GetTeacher(id int) (*model.TeacherDB, error) {
	teacher := &model.TeacherDB{}
	rows, err := us.db.Query(context.Background(), `SELECT id, login, name, tgAccount, vkAccount, tgBotLink, vkBotLink FROM teachers WHERE id  = $1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&teacher.ID, &teacher.Login, &teacher.Name, &teacher.TgAccount, &teacher.VkAccount, &teacher.TgBotLink, &teacher.VkBotLink)
		if err != nil {
			return nil, err
		}
	}
	return teacher, nil
}

func (us *Store) UpdateTeacher(in *model.TeacherDB) error {
	return nil
}

func (us *Store) AddStudent(in *model.StudentDB) error {
	_, err := us.db.Query(context.Background(), `INSERT INTO students (inviteHash, name) VALUES ($1, $2);`, in.InviteHash, in.Name)
	if err != nil {
		return err
	}
	return nil
}

func (us *Store) UpdateStudent(in *model.StudentDB) error {
	return nil
}

func (us *Store) CreateChat(in *model.ChatDB) error {
	_, err := us.db.Query(context.Background(), `INSERT INTO chats (teacherID, studentHash) VALUES ($1, $2);`, in.TeacherID, in.StudentHash)
	if err != nil {
		return err
	}
	return nil
}

func (us *Store) AddMessage(in *model.CreateMessage) error {
	_, err := us.db.Query(context.Background(), `INSERT INTO messages (chatID, text, isAuthorTeacher, time) VALUES ($1, $2, $3, $5);`, in.ChatID, in.Text, in.IsAuthorTeacher, time.Now().Format("2006.01.02 15:04:05"))
	if err != nil {
		return err
	}
	return nil
}

// func (us *Store) GetChatID(in *model.ChatDB) (int, error) {
// 	id := 0
// 	rows, err := us.db.Query(`SELECT id FROM chats WHERE teacherID  = $1 AND studentHash = $2`, in.TeacherID, in.StudentHash)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		err := rows.Scan(&id)
// 		if err != nil {
// 			return 0, err
// 		}

// 	}
// 	return id, nil
// }

func (us *Store) GetChatFromDB(id int) (*model.Chat, error) {
	// messages := []*model.MessageDB{}
	// chatID, err := us.GetChatID(in)
	// if err != nil || chatID == 0 {
	// 	return nil, err
	// }
	messages := []*model.MessageChat{}
	rows, err := us.db.Query(context.Background(), `SELECT text, isAuthorTeacher, attaches, time FROM messages WHERE chatID  = $1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		dat := model.MessageChat{}
		err := rows.Scan(&dat.Text, &dat.IsAuthorTeacher, &dat.Attaches, &dat.Time)
		if err != nil {
			return nil, err
		}
		//mes := model.MessageChat{Text: dat.Text, IsAuthorTeacher: dat.IsAuthorTeacher, }
		messages = append(messages, &dat)
	}

	return &model.Chat{Messages: messages}, nil
}

func (us *Store) GetChatsByID(idTeacher int) (*model.Chats, error) {
	//chats := []*model.Chat{}
	chats := []*model.ChatInfo{}
	rows, err := us.db.Query(context.Background(), `SELECT id FROM chats WHERE teacherID = $1`, idTeacher)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		tmpID := 1
		err := rows.Scan(&tmpID)
		if err != nil {
			return nil, err
		}
		//chatsIDs = append(chatsIDs, &dat)
		// chat, err := us.GetChatFromDB(tmpID)
		// if err != nil {
		// 	return nil, err
		// }
		tmpChat := &model.ChatInfo{ChatID: tmpID}
		chats = append(chats, tmpChat)
	}

	return &model.Chats{Chats: chats}, nil
}
