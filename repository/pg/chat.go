package pg

import (
	e "main/domain/errors"
	m "main/domain/model"

	"github.com/lib/pq"
)

func (s *PostgreSqlStore) CheckChatExistence(id int) error {
	var tmp int
	if err := s.db.QueryRow(
		`SELECT 1 FROM chats WHERE id = $1;`,
		id,
	).Scan(&tmp); err != nil {
		return e.StacktraceError(err)
	}
	return nil
}

func (s *PostgreSqlStore) GetChatByID(id int) (*m.Chat, error) {
	rows, err := s.db.Query(
		`SELECT id, text, isAuthorTeacher, attaches, createTime, isRead FROM messages
		 WHERE chatID = $1;`,
		id,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	messages := []m.Message{}
	for rows.Next() {
		var tmpMsg m.Message

		if err := rows.Scan(
			&tmpMsg.ID, &tmpMsg.Text,
			&tmpMsg.IsAuthorTeacher, (*pq.StringArray)(&tmpMsg.Attaches),
			&tmpMsg.CreateTime, &tmpMsg.IsRead,
		); err != nil {
			return nil, e.StacktraceError(err)
		}

		messages = append(messages, tmpMsg)
	}
	return &m.Chat{Messages: messages}, nil
}

func (s *PostgreSqlStore) ReadChatByID(id int, teacherID int) error {
	var tID int
	if err := s.db.QueryRow(
		`SELECT teacherID FROM chats WHERE id = $1;`,
		id,
	).Scan(&tID); err != nil {
		return e.StacktraceError(err)
	}
	if tID != teacherID {
		return e.StacktraceError(e.ErrForbidden403)
	}
	_, err := s.db.Exec(
		`UPDATE messages set isRead = TRUE WHERE chatID = $1;`, id)
	if err != nil {
		return e.StacktraceError(err)
	}
	return nil
}

func (s *PostgreSqlStore) GetChatsByTeacherID(teacherID int) ([]m.ChatPreview, error) {
	rows, err := s.db.Query(
		`SELECT m1.chatID, s.id, s.name, s.socialType, s.avatar, m1.text, m1.createTime, m1.isRead
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

	chats := []m.ChatPreview{}
	for rows.Next() {
		tmpChat := m.ChatPreview{}

		if err = rows.Scan(
			&tmpChat.ChatID,
			&tmpChat.StudentID,
			&tmpChat.StudentName,
			&tmpChat.SocialType,
			&tmpChat.StudentAvatar,
			&tmpChat.LastMessageText,
			&tmpChat.LastMessageDate,
			&tmpChat.IsRead,
		); err != nil {
			return nil, e.StacktraceError(err)
		}
		chats = append(chats, tmpChat)
	}

	return chats, nil
}

func (s *PostgreSqlStore) GetChatIDBySolutionID(solutionID int) (int, error) {
	var id int
	if err := s.db.QueryRow(
		`SELECT c.id FROM chats c
		 JOIN solutions s ON c.studentID = s.studentID
		 WHERE s.id = $1`,
		solutionID,
	).Scan(&id); err != nil {
		return 0, e.StacktraceError(err)
	}

	return id, nil
}
