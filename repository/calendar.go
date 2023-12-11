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

func (s *Store) CreateEventDB(in *model.CalendarEvent) error {
	_, err := s.db.Exec(`INSERT INTO events (id, classID, title, description, startDate, endDate) VALUES ($1, $2, $3, $4, $5, $6)`, in.ID, in.ClassID, in.Title, in.Description, in.StartDate, in.EndDate)
	if err != nil {
		return e.StacktraceError(err)
	}
	return nil
}

func (s *Store) DeleteEventDB(id string) error {
	_, err := s.db.Exec(`DELETE FROM events WHERE id = $1;`, id)
	if err != nil {
		return e.StacktraceError(err)
	}
	return nil
}
