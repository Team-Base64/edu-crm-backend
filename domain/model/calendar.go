package model

type OAUTH2Token struct {
	Token string `json:"token"`
}

type CreateCalendarResponse struct {
	ID         int    `json:"id"`
	IDInGoogle string `json:"googleid"`
}

type CreateCalendarEvent struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
}
