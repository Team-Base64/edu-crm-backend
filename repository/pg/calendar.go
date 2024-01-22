package pg

import (
	e "main/domain/errors"
	m "main/domain/model"
)

func (s *PostgreSqlStore) GetCalendar(teacherID int) (*m.CalendarParams, error) {
	ans := m.CalendarParams{}
	row := s.db.QueryRow(`SELECT id, internalApiID FROM calendars WHERE teacherID = $1;`, teacherID)
	if err := row.Scan(&ans.ID, &ans.InternalApiID); err != nil {
		return nil, e.StacktraceError(err)
	}
	return &ans, nil
}
