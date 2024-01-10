package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	e "main/domain/errors"
	m "main/domain/model"
)

// GetChats godoc
// @Summary Get chats of teacher
// @Description Get chats of teacher
// @ID getChats
// @Accept  json
// @Produce  json
// @Tags Chats
// @Success 200 {object} m.ChatPreviewList
// @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} m.Error "internal server error - Request is valid but operation failed at server side"
// @Router /chats [get]
func (api *Handler) GetTeacherChats(w http.ResponseWriter, r *http.Request) {
	teacherProfile := r.Context().Value(KeyUserdata{"userdata"}).(*m.TeacherDB)
	chats, err := api.usecase.GetChatsByTeacherID(teacherProfile.ID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(&m.ChatPreviewList{Chats: chats})
}

// GetChat godoc
// @Summary Get chat messages by id
// @Description Get chats messages by chat id
// @ID getChat
// @Accept  json
// @Produce  json
// @Tags Chats
// @Param chatID path string true "Chat id"
// @Success 200 {object} m.Chat
// @Failure 400 {object} m.Error "bad request - Problem with the request"
// @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} m.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} m.Error "internal server error - Request is valid but operation failed at server side"
// @Router /chats/{chatID} [get]
func (api *Handler) GetChat(w http.ResponseWriter, r *http.Request) {
	s := strings.Split(r.URL.Path, "/")
	idS := s[len(s)-1]
	id, err := strconv.Atoi(idS)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	chat, err := api.usecase.GetChatByID(id)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(chat)
}

// ReadChat godoc
// @Summary Read chat messages by id
// @Description Read chats messages by chat id
// @ID readChat
// @Accept  json
// @Produce  json
// @Tags Chats
// @Param chatID path string true "Chat id"
// @Success 200 {object} m.Response "OK"
// @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// @Failure 403 {object} m.Error "forbidden"
// @Failure 404 {object} m.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} m.Error "internal server error - Request is valid but operation failed at server side"
// @Router /chats/{chatID} [post]
func (api *Handler) ReadChat(w http.ResponseWriter, r *http.Request) {
	teacherProfile := r.Context().Value(KeyUserdata{"userdata"}).(*m.TeacherDB)
	s := strings.Split(r.URL.Path, "/")
	idS := s[len(s)-1]
	id, err := strconv.Atoi(idS)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	err = api.usecase.ReadChatByID(id, teacherProfile.ID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(&m.Response{})
}
