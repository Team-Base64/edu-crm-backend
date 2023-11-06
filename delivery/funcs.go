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
// @Success 200 {object} model.ClassesInfo
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /classes [get]
func (api *Handler) GetTeacherClasses(w http.ResponseWriter, r *http.Request) {
	mockTeacherID := 1

	classes, err := api.usecase.GetClassesByTeacherID(mockTeacherID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, err)
		return
	}
	json.NewEncoder(w).Encode(classes)
}

// GetClasses godoc
// @Summary Get class by id
// @Description Get class by id
// @ID getClass
// @Accept  json
// @Produce  json
// @Param classID path string true "Class id"
// @Success 200 {object} model.ClassInfo
// @Failure 400 {object} model.Error "bad request - Problem with the request"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} model.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /classes/{classID} [get]
func (api *Handler) GetClass(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	classID, err := strconv.Atoi(path[len(path)-1])
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, e.ErrBadRequest400)
	}

	class, err := api.usecase.GetClassByID(classID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, err)
		return
	}
	json.NewEncoder(w).Encode(class)
}

// CreateClass godoc
// @Summary Create class
// @Description Create class
// @ID createClass
// @Accept  json
// @Produce  json
// @Param class body model.ClassCreate true "Class for creating"
// @Success 200 {object} model.ClassCreateResponse
// @Failure 400 {object} model.Error "bad request - Problem with the request"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} model.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /classes [post]
func (api *Handler) CreateClass(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newClass model.ClassCreate
	err := decoder.Decode(&newClass)
	if err != nil {
		ReturnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	res, err := api.usecase.CreateClass(mockTeacherID, &newClass)
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, err)
		return
	}
	json.NewEncoder(w).Encode(res)
}

// GetStudentsFromClass godoc
// @Summary Get students from class
// @Description Get students from class by class id
// @ID getStudentsFromClass
// @Accept  json
// @Produce  json
// @Param classID path string true "Class id"
// @Success 200 {object} model.StudentsFromClass
// @Failure 400 {object} model.Error "bad request - Problem with the request"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} model.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /classes/{classID}/students [get]
func (api *Handler) GetStudentsFromClass(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	chatId, err := strconv.Atoi(path[len(path)-2])
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	students, err := api.usecase.GetStudentsFromClass(chatId)
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(students)
}

// GetClassFeed godoc
// @Summary Get class feed
// @Description Get posts from class by class id
// @ID getClassFeed
// @Accept  json
// @Produce  json
// @Param classID path string true "Class id"
// @Success 200 {object} model.Feed
// @Failure 400 {object} model.Error "bad request - Problem with the request"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} model.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /classes/{classID}/feed [get]
func (api *Handler) GetClassFeed(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	chatId, err := strconv.Atoi(path[len(path)-2])
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	feed, err := api.usecase.GetClassFeed(chatId)
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(feed)
}

// GetHomeworksFromClass godoc
// @Summary Get class homeworks
// @Description Get homeworks from class by class id
// @ID getHomeworksFromClass
// @Accept  json
// @Produce  json
// @Param classID path string true "Class id"
// @Success 200 {object} model.HomeworksFromClass
// @Failure 400 {object} model.Error "bad request - Problem with the request"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} model.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /classes/{classID}/homeworks [get]
func (api *Handler) GetHomeworksFromClass(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	chatId, err := strconv.Atoi(path[len(path)-2])
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	hws, err := api.usecase.GetHomeworksByClassID(chatId)
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(hws)
}

// GetHomework godoc
// @Summary Get homework
// @Description Get homework by id
// @ID getHomework
// @Accept  json
// @Produce  json
// @Param hwID path string true "Homework id"
// @Success 200 {object} model.HomeworkByID
// @Failure 400 {object} model.Error "bad request - Problem with the request"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} model.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /homeworks/{hwID} [get]
func (api *Handler) GetHomework(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	hwID, err := strconv.Atoi(path[len(path)-1])
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	hw, err := api.usecase.GetHomeworkByID(hwID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(hw)
}
