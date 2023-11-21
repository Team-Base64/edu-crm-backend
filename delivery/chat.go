package delivery

import (
	"encoding/json"
	"log"
	e "main/domain/errors"
	"main/domain/model"
	"net/http"
	"strconv"
	"strings"
)

// GetChats godoc
// @Summary Get chats of teacher
// @Description Get chats of teacher
// @ID getChats
// @Accept  json
// @Produce  json
// @Tags Chats
// @Success 200 {object} model.ChatPreviewList
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /chats [get]
func (api *Handler) GetTeacherChats(w http.ResponseWriter, r *http.Request) {
	teacherProfile := r.Context().Value(KeyUserdata{"userdata"}).(*model.TeacherDB)
	chats, err := api.usecase.GetChatsByTeacherID(teacherProfile.ID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(&model.ChatPreviewList{Chats: chats})
}

// GetChat godoc
// @Summary Get chat messages by id
// @Description Get chats messages by chat id
// @ID getChat
// @Accept  json
// @Produce  json
// @Tags Chats
// @Param chatID path string true "Chat id"
// @Success 200 {object} model.Chat
// @Failure 400 {object} model.Error "bad request - Problem with the request"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} model.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
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
