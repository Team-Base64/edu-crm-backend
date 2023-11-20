package model

type OAUTH2Token struct {
	Token string `json:"token"`
}

type CreateCalendarResponse struct {
	ID         int    `json:"id"`
	IDInGoogle string `json:"googleid"`
}

type CalendarEvent struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	ClassID     int    `json:"classid"`
	ID          string `json:"id,omitempty"`
}

type CalendarEvents struct {
	Events []CalendarEvent `json:"events"`
}

type DeleteEvent struct {
	ID string `json:"id"`
}
