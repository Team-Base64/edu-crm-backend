package repository

import (
	e "main/domain/errors"
	"main/domain/model"
)

func (s *Store) GetTokenDB(id int) (string, error) {
	var tok string
	row := s.db.QueryRow(
		`SELECT OAUTH2Token FROM teachers WHERE id = $1;`, id)
	if err := row.Scan(&tok); err != nil {
		return "", e.StacktraceError(err)
	}
	return tok, nil
}

func (s *Store) GetCalendarDB(teacherID int) (*model.CalendarParams, error) {
	// tok, err := s.GetTokenDB(teacherID)
	// if err != nil {
	// 	return 0, e.StacktraceError(err)
	// }
	ans := model.CalendarParams{}
	row := s.db.QueryRow(`SELECT id, idInGoogle FROM calendars WHERE teacherID = $1;`, teacherID)
	if err := row.Scan(&ans.ID, &ans.IDInGoogle); err != nil {
		return nil, e.StacktraceError(err)
	}
	return &ans, nil
}
