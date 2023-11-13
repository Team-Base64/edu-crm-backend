package repository

import (
	e "main/domain/errors"
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

func (s *Store) CreateCalendarDB(teacherID int, googleID string) (int, error) {
	// tok, err := s.GetTokenDB(teacherID)
	// if err != nil {
	// 	return 0, e.StacktraceError(err)
	// }
	id := 1
	err := s.db.QueryRow(`INSERT INTO calendars (teacherID, idInGoogle) VALUES ($1, $2) RETURNING id;`, teacherID, googleID).Scan(&id)
	if err != nil {
		return 0, e.StacktraceError(err)
	}
	return id, nil
}

func (s *Store) GetCalendarGoogleID(teacherID int) (string, error) {
	// tok, err := s.GetTokenDB(teacherID)
	// if err != nil {
	// 	return 0, e.StacktraceError(err)
	// }
	var id string
	row := s.db.QueryRow(`SELECT idInGoogle FROM calendars WHERE teacherID = $1;`, teacherID)
	if err := row.Scan(&id); err != nil {
		return "", e.StacktraceError(err)
	}
	return id, nil
}
