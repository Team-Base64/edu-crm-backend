package repository

import (
	e "main/domain/errors"
	"main/domain/model"

	"github.com/lib/pq"
)

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
