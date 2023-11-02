package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"main/domain/model"

	e "main/domain/errors"
	uc "main/usecase"
)

var mockTeacherID = 1

// @title TCRA API
// @version 1.0
// @description TCRA back server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath  /api

type Handler struct {
	usecase uc.UsecaseInterface
}

func NewHandler(uc uc.UsecaseInterface) *Handler {
	return &Handler{
		usecase: uc,
	}
}

func ReturnErrorJSON(w http.ResponseWriter, err error) {
	errCode, errText := e.CheckError(err)
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(&model.Error{Error: errText})
}

// CreateTeacher godoc
// @Summary Create teacher
// @Description Create teacher
// @ID createTeacher
// @Accept  json
// @Produce  json
// @Param user body model.TeacherSignUp true "Teacher params"
// @Success 200 {object} model.Response "OK"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /register [post]
func (api *Handler) CreateTeacher(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req model.TeacherSignUp
	err := decoder.Decode(&req)
	if err != nil {
		ReturnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	err = api.usecase.CreateTeacher(&req)
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(&model.Response{})
}

// GetTeacher godoc
// @Summary Get teacher's info
// @Description gets teacher's info
// @ID getTeacher
// @Accept  json
// @Produce  json
// @Success 200 {object} model.TeacherProfile
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /profile [get]
func (api *Handler) GetTeacherProfile(w http.ResponseWriter, r *http.Request) {
	teacher, err := api.usecase.GetTeacherProfile(mockTeacherID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(teacher)
}

// GetChats godoc
// @Summary Get chats of teacher
// @Description Get chats of teacher
// @ID getChats
// @Accept  json
// @Produce  json
// @Success 200 {object} model.ChatPreviewList
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /chats [get]
func (api *Handler) GetTeacherChats(w http.ResponseWriter, r *http.Request) {
	chats, err := api.usecase.GetChatsByTeacherID(mockTeacherID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(chats)
}

// GetChat godoc
// @Summary Get chat messages by id
// @Description Get chats messages by chat id
// @ID getChat
// @Accept  json
// @Produce  json
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
		ReturnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	msgs, err := api.usecase.GetChatByID(id)
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(msgs)
}

// GetClasses godoc
// @Summary Get teacher`s classes
// @Description Get teacher`s classes
// @ID getClasses
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Classes
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /class [get]
func (api *Handler) GetTeacherClasses(w http.ResponseWriter, r *http.Request) {
	mockTeacherID := 1

	classes, err := api.usecase.GetClassesByTeacherID(mockTeacherID)
	if err != nil {
		log.Println("usecase: ", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	json.NewEncoder(w).Encode(classes)
}
