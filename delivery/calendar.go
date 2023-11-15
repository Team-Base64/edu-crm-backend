package delivery

import (
	"encoding/json"
	"log"
	e "main/domain/errors"
	"main/domain/model"
	"net/http"
)

// SetOAUTH2Token godoc
// @Summary Sets teacher's OAUTH2Token
// @Description Sets teacher's OAUTH2Token
// @ID SetOAUTH2Token
// @Accept  json
// @Produce  json
// @Tags Calendar
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /oauth [post]
func (api *Handler) SetOAUTH2Token(w http.ResponseWriter, r *http.Request) {
	err := api.usecase.SetOAUTH2Token()
	if err != nil {
		log.Println("Error ", err)
		returnErrorJSON(w, e.ErrServerError500)
		return
	}
	json.NewEncoder(w).Encode(&model.Response{})
}

// SaveOAUTH2TokenToFile godoc
// @Summary Saves teacher's OAUTH2Token
// @Description Saves teacher's OAUTH2Token
// @ID SaveOAUTH2TokenToFile
// @Accept  json
// @Produce  json
// @Tags Calendar
// @Param   code         query     string  true  "code"
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /oauth/savetoken [get]
func (api *Handler) SaveOAUTH2TokenToFile(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	err := api.usecase.SaveOAUTH2Token(code)
	if err != nil {
		log.Println("Error ", err)
		returnErrorJSON(w, e.ErrServerError500)
		return
	}
	json.NewEncoder(w).Encode(&model.Response{})
}

// CreateCalendar godoc
// @Summary Creates teacher's calendar
// @Description Creates teacher's calendar
// @ID CreateCalendar
// @Accept  json
// @Produce  json
// @Tags Calendar
// @Success 200 {object} model.CreateCalendarResponse
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /calendar [post]
func (api *Handler) CreateCalendar(w http.ResponseWriter, r *http.Request) {
	mockTeackerID := 1
	createdResponse, err := api.usecase.CreateCalendar(mockTeackerID)
	if err != nil {
		log.Println("Error ", err)
		returnErrorJSON(w, e.ErrServerError500)
		return
	}

	json.NewEncoder(w).Encode(createdResponse)
}

// CreateCalendarEvent godoc
// @Summary Creates teacher's calendar event
// @Description Creates teacher's calendar event
// @ID CreateCalendarEvent
// @Accept  json
// @Produce  json
// @Tags Calendar
// @Param event body model.CalendarEvent true "Event for creating"
// @Success 200 {object} model.Response
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /calendar/addevent [post]
func (api *Handler) CreateCalendarEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	mockTeackerID := 1

	decoder := json.NewDecoder(r.Body)
	var req model.CalendarEvent
	if err := decoder.Decode(&req); err != nil {
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	err := api.usecase.CreateCalendarEvent(&req, mockTeackerID)
	if err != nil {
		log.Println("Error ", err)
		returnErrorJSON(w, e.ErrServerError500)
		return
	}

	json.NewEncoder(w).Encode(&model.Response{})
}

// GetCalendarEvents godoc
// @Summary Get teacher's calendar event
// @Description Get teacher's calendar event
// @ID GetCalendarEvents
// @Accept  json
// @Produce  json
// @Tags Calendar
// @Success 200 {object} model.CalendarEvent
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /calendar/events [get]
func (api *Handler) GetCalendarEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	mockTeacherID := 1

	events, err := api.usecase.GetCalendarEvents(mockTeacherID)
	if err != nil {
		log.Println("Error ", err)
		returnErrorJSON(w, e.ErrServerError500)
		return
	}

	json.NewEncoder(w).Encode(&model.CalendarEvents{Events: events})
}

// DeleteCalendarEvent godoc
// @Summary Delete teacher's calendar event
// @Description Delete teacher's calendar event
// @ID DeleteCalendarEvent
// @Accept  json
// @Produce  json
// @Tags Calendar
// @Param event body model.DeleteEvent true "Event for deleting"
// @Success 200 {object} model.Response
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /calendar/event [delete]
func (api *Handler) DeleteCalendarEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	mockTeacherID := 1

	decoder := json.NewDecoder(r.Body)
	var req model.DeleteEvent
	if err := decoder.Decode(&req); err != nil {
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	err := api.usecase.DeleteCalendarEvent(mockTeacherID, req.ID)
	if err != nil {
		log.Println("Error ", err)
		returnErrorJSON(w, e.ErrServerError500)
		return
	}

	json.NewEncoder(w).Encode(&model.Response{})
}

// UpdateCalendarEvent godoc
// @Summary Update teacher's calendar event
// @Description Update teacher's calendar event
// @ID UpdateCalendarEvent
// @Accept  json
// @Produce  json
// @Tags Calendar
// @Param event body model.CalendarEvent true "Event for updating"
// @Success 200 {object} model.Response
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /calendar/event [post]
func (api *Handler) UpdateCalendarEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	mockTeacherID := 1

	decoder := json.NewDecoder(r.Body)
	var req model.CalendarEvent
	if err := decoder.Decode(&req); err != nil {
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}
	if req.ID == "" {
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	err := api.usecase.UpdateCalendarEvent(&req, mockTeacherID)
	if err != nil {
		log.Println("Error ", err)
		returnErrorJSON(w, e.ErrServerError500)
		return
	}

	json.NewEncoder(w).Encode(&model.Response{})
}
