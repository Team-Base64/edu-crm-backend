package delivery

import (
	"encoding/json"
	"log"
	baseErrors "main/domain/errors"
	"main/domain/model"
	"net/http"
	"strconv"
	"strings"

	usecase "main/usecase"
)

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
	usecase usecase.UsecaseInterface
}

func NewHandler(uc usecase.UsecaseInterface) *Handler {
	return &Handler{
		usecase: uc,
	}
}

func ReturnErrorJSON(w http.ResponseWriter, err error, errCode int) {
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(&model.Error{Error: err.Error()})
	return
}

// CreateTeacher godoc
// @Summary Create teacher
// @Description Create teacher
// @ID createTeacher
// @Accept  json
// @Produce  json
// @Param user body model.TeacherDB true "Teacher params"
// @Success 200 {object} model.Response "OK"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /register [post]
func (api *Handler) CreateTeacher(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req model.TeacherDB
	err := decoder.Decode(&req)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	err = api.usecase.CreateTeacher(&req)
	if err != nil {
		log.Println("err", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
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
// @Success 200 {object} model.TeacherDB
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /profile [get]
func (api *Handler) GetTeacher(w http.ResponseWriter, r *http.Request) {

	teacher, err := api.usecase.GetTeacher(1)
	if err != nil {
		log.Println("err", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	json.NewEncoder(w).Encode(teacher)
}

// ChangeUser godoc
// @Summary changes teacher's parameters
// @Description changes teacher's parameters
// @ID changeUserParameters
// @Accept  json
// @Produce  json
// @Param user body model.TeacherDB true "Teacher params"
// @Success 200 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /profile [post]
func (api *Handler) ChangeProfile(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req model.TeacherDB
	err := decoder.Decode(&req)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	err = api.usecase.ChangeTeacher(&req)
	if err != nil {
		log.Println("err", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	json.NewEncoder(w).Encode(&model.Response{})
}

// AddStudent godoc
// @Summary Add student
// @Description Add student
// @ID addStudent
// @Accept  json
// @Produce  json
// @Param user body model.CreateStudentDB true "Student params"
// @Success 200 {object} model.Response "OK"
// @Failure 400 {object} model.Error "Bad request - Problem with the request"
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /addstudent [post]
func (api *Handler) AddStudent(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req model.CreateStudentDB
	err := decoder.Decode(&req)
	if err != nil {
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	err = api.usecase.AddStudent(&req)
	if err != nil {
		log.Println("err", err)
		ReturnErrorJSON(w, baseErrors.ErrServerError500, 500)
		return
	}
	json.NewEncoder(w).Encode(&model.Response{})
}

// GetChats godoc
// @Summary Get chats of teacher
// @Description Get chats of teacher
// @ID getChats
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Chats
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /chats [get]
func (api *Handler) GetTeacherChats(w http.ResponseWriter, r *http.Request) {
	// s := strings.Split(r.URL.Path, "/")
	// idS := s[len(s)-1]
	// id, err := strconv.Atoi(idS)
	// if err != nil {
	// 	log.Println("error: ", err)
	// 	ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
	// 	return
	// }
	mockTeacherID := 1
	chats, err := api.usecase.GetChatsByTeacherID(mockTeacherID)
	if err != nil {
		log.Println("err", err)
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 500)
		return
	}
	json.NewEncoder(w).Encode(chats)
}

// GetChat godoc
// @Summary Get chat by id
// @Description Get chats messages by chat id
// @ID getChat
// @Accept  json
// @Produce  json
// @Param chatID path string true "Chat id"
// @Success 200 {object} model.Chat
// @Failure 401 {object} model.Error "Unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "Internal Server Error - Request is valid but operation failed at server side"
// @Router /chats/{chatID} [get]
func (api *Handler) GetChat(w http.ResponseWriter, r *http.Request) {
	s := strings.Split(r.URL.Path, "/")
	idS := s[len(s)-1]
	id, err := strconv.Atoi(idS)
	if err != nil {
		log.Println("error: ", err)
		ReturnErrorJSON(w, baseErrors.ErrBadRequest400, 400)
		return
	}

	chats, err := api.usecase.GetChatByID(id)
	if err != nil {
		log.Println("err", err)
		ReturnErrorJSON(w, baseErrors.ErrUnauthorized401, 500)
		return
	}
	json.NewEncoder(w).Encode(chats)
}
