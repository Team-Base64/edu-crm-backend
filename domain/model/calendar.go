package model

import "time"

type CalendarParams struct {
	ID            int    `json:"id"`
	InternalApiID string `json:"googleid"`
}

type CalendarEvent struct {
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	ClassID     int       `json:"classid"`
	ID          string    `json:"id,omitempty"`
}

type CreateEvent struct {
	CalendarID string         `json:"calendarid"`
	Event      *CalendarEvent `json:"event"`
}

type CalendarEvents struct {
	Events []CalendarEvent `json:"events"`
}

type DeleteEvent struct {
	ID string `json:"id"`
}
