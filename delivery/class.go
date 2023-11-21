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

// GetClasses godoc
// @Summary Get teacher`s classes
// @Description Get teacher`s classes
// @ID getClasses
// @Accept  json
// @Produce  json
// @Tags Class
// @Success 200 {object} model.ClassInfoList
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /classes [get]
func (api *Handler) GetTeacherClasses(w http.ResponseWriter, r *http.Request) {
	teacherProfile := r.Context().Value(KeyUserdata{"userdata"}).(*model.TeacherDB)
	classes, err := api.usecase.GetClassesByTeacherID(teacherProfile.ID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
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
// @Tags Class
// @Param classID path string true "Class id"
// @Success 200 {object} model.ClassInfoResponse
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
		returnErrorJSON(w, e.ErrBadRequest400)
	}

	class, err := api.usecase.GetClassByID(classID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}
	json.NewEncoder(w).Encode(&model.ClassInfoResponse{Class: *class})
}

// CreateClass godoc
// @Summary Create class
// @Description Create class
// @ID createClass
// @Accept  json
// @Produce  json
// @Tags Class
// @Param class body model.ClassCreate true "Class for creating"
// @Success 200 {object} model.ClassInfoResponse
// @Failure 400 {object} model.Error "bad request - Problem with the request"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} model.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /classes [post]
func (api *Handler) CreateClass(w http.ResponseWriter, r *http.Request) {
	teacherProfile := r.Context().Value(KeyUserdata{"userdata"}).(*model.TeacherDB)
	decoder := json.NewDecoder(r.Body)
	var newClass model.ClassCreate
	if err := decoder.Decode(&newClass); err != nil {
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	class, err := api.usecase.CreateClass(teacherProfile.ID, &newClass)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(&model.ClassInfoResponse{Class: *class})
}
