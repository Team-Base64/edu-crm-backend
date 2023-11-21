package model

import "time"

type OAUTH2Token struct {
	Token string `json:"token"`
}

type CalendarParams struct {
	ID         int    `json:"id"`
	IDInGoogle string `json:"googleid"`
}

type CalendarEvent struct {
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	ClassID     int       `json:"classid"`
	ID          string    `json:"id,omitempty"`
}

type CalendarEvents struct {
	Events []CalendarEvent `json:"events"`
}

type DeleteEvent struct {
	ID string `json:"id"`
}
