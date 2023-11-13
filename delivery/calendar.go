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
	//decoder := json.NewDecoder(r.Body)
	// var tok model.OAUTH2Token
	// if err := decoder.Decode(&tok); err != nil {
	// 	returnErrorJSON(w, e.ErrBadRequest400)
	// 	return
	// }
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
	//decoder := json.NewDecoder(r.Body)
	// var tok model.OAUTH2Token
	// if err := decoder.Decode(&tok); err != nil {
	// 	returnErrorJSON(w, e.ErrBadRequest400)
	// 	return
	// }
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
// @Success 200 {object} model.Response
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /calendar/addevent [post]
func (api *Handler) CreateCalendarEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	mockTeackerID := 1
	mockClassID := 1

	decoder := json.NewDecoder(r.Body)
	var req model.CreateCalendarEvent
	if err := decoder.Decode(&req); err != nil {
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	err := api.usecase.CreateCalendarEvent(&req, mockTeackerID, mockClassID)
	if err != nil {
		log.Println("Error ", err)
		returnErrorJSON(w, e.ErrServerError500)
		return
	}

	json.NewEncoder(w).Encode(&model.Response{})
}
