package delivery

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	e "main/domain/errors"
	m "main/domain/model"

	"github.com/microcosm-cc/bluemonday"
)

// CreateCalendar godoc
// @Summary Creates teacher's calendar
// @Description Creates teacher's calendar
// @ID CreateCalendar
// @Accept  json
// @Produce  json
// @Tags Calendar
// @Success 200 {object} m.CalendarParams
// @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} m.Error "internal server error - Request is valid but operation failed at server side"
// @Router /calendar [post]
func (api *Handler) CreateCalendar(w http.ResponseWriter, r *http.Request) {
	teacherProfile := r.Context().Value(KeyUserdata{"userdata"}).(*m.TeacherDB)
	err := api.usecase.CreateCalendar(teacherProfile.ID)
	if err != nil {
		log.Println("Error ", err)
		returnErrorJSON(w, e.ErrServerError500)
		return
	}

	//json.NewEncoder(w).Encode(createdResponse)
	json.NewEncoder(w).Encode(&m.Response{})
}

// GetCalendar godoc
// @Summary Gets teacher's calendar
// @Description Gets teacher's calendar
// @ID GetCalendar
// @Accept  json
// @Produce  json
// @Tags Calendar
// @Success 200 {object} m.CalendarParams
// @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} m.Error "internal server error - Request is valid but operation failed at server side"
// @Router /calendar [get]
func (api *Handler) GetCalendar(w http.ResponseWriter, r *http.Request) {
	teacherProfile := r.Context().Value(KeyUserdata{"userdata"}).(*m.TeacherDB)
	createdResponse, err := api.usecase.GetCalendarParams(teacherProfile.ID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrServerError500)
		return
	}

	json.NewEncoder(w).Encode(createdResponse)
	json.NewEncoder(w).Encode(&m.Response{})
}

// CreateCalendarEvent godoc
// @Summary Creates teacher's calendar event
// @Description Creates teacher's calendar event
// @ID CreateCalendarEvent
// @Accept  json
// @Produce  json
// @Tags Calendar
// @Param event body m.CalendarEvent true "Event for creating"
// @Success 200 {object} m.Response
// @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} m.Error "internal server error - Request is valid but operation failed at server side"
// @Router /calendar/addevent [post]
func (api *Handler) CreateCalendarEvent(w http.ResponseWriter, r *http.Request) {
	teacherProfile := r.Context().Value(KeyUserdata{"userdata"}).(*m.TeacherDB)
	decoder := json.NewDecoder(r.Body)
	var req m.CalendarEvent
	if err := decoder.Decode(&req); err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}
	sanitizer := bluemonday.UGCPolicy()
	req.Title = sanitizer.Sanitize(req.Title)
	req.Description = sanitizer.Sanitize(req.Description)
	err := api.usecase.CreateCalendarEvent(&req, teacherProfile.ID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrServerError500)
		return
	}

	json.NewEncoder(w).Encode(&m.Response{})
}

// GetCalendarEvents godoc
// @Summary Get teacher's calendar event
// @Description Get teacher's calendar event
// @ID GetCalendarEvents
// @Accept  json
// @Produce  json
// @Tags Calendar
// @Success 200 {object} m.CalendarEvents
// @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} m.Error "internal server error - Request is valid but operation failed at server side"
// @Router /calendar/events [get]
func (api *Handler) GetCalendarEvents(w http.ResponseWriter, r *http.Request) {
	teacherProfile := r.Context().Value(KeyUserdata{"userdata"}).(*m.TeacherDB)
	events, err := api.usecase.GetCalendarEvents(teacherProfile.ID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrServerError500)
		return
	}

	json.NewEncoder(w).Encode(&m.CalendarEvents{Events: events})
}

// DeleteCalendarEvent godoc
// @Summary Delete teacher's calendar event
// @Description Delete teacher's calendar event
// @ID DeleteCalendarEvent
// @Accept  json
// @Produce  json
// @Tags Calendar
// @Param event body m.DeleteEvent true "Event for deleting"
// @Success 200 {object} m.Response
// @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} m.Error "internal server error - Request is valid but operation failed at server side"
// @Router /calendar/event [delete]
func (api *Handler) DeleteCalendarEvent(w http.ResponseWriter, r *http.Request) {
	teacherProfile := r.Context().Value(KeyUserdata{"userdata"}).(*m.TeacherDB)
	decoder := json.NewDecoder(r.Body)
	var req m.DeleteEvent
	if err := decoder.Decode(&req); err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	err := api.usecase.DeleteCalendarEvent(teacherProfile.ID, req.ID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrServerError500)
		return
	}

	json.NewEncoder(w).Encode(&m.Response{})
}

// UpdateCalendarEvent godoc
// @Summary Update teacher's calendar event
// @Description Update teacher's calendar event
// @ID UpdateCalendarEvent
// @Accept  json
// @Produce  json
// @Tags Calendar
// @Param event body m.CalendarEvent true "Event for updating"
// @Success 200 {object} m.Response
// @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} m.Error "internal server error - Request is valid but operation failed at server side"
// @Router /calendar/event [post]
func (api *Handler) UpdateCalendarEvent(w http.ResponseWriter, r *http.Request) {
	teacherProfile := r.Context().Value(KeyUserdata{"userdata"}).(*m.TeacherDB)
	decoder := json.NewDecoder(r.Body)
	var req m.CalendarEvent
	if err := decoder.Decode(&req); err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}
	if req.ID == "" {
		log.Println(e.StacktraceError(errors.New("empty ID")))
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	sanitizer := bluemonday.UGCPolicy()
	req.Title = sanitizer.Sanitize(req.Title)
	req.Description = sanitizer.Sanitize(req.Description)
	err := api.usecase.UpdateCalendarEvent(&req, teacherProfile.ID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrServerError500)
		return
	}

	json.NewEncoder(w).Encode(&m.Response{})
}
