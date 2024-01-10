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

// GetStudent godoc
// @Summary Get student
// @Description Get student by id
// @ID getStudent
// @Accept  json
// @Produce  json
// @Tags Student
// @Param studentID path string true "StudentID id"
// @Success 200 {object} m.StudentByIDResponse
// @Failure 400 {object} m.Error "bad request - Problem with the request"
// @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} m.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} m.Error "internal server error - Request is valid but operation failed at server side"
// @Router /students/{studentID} [get]
func (api *Handler) GetStudent(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	studentID, err := strconv.Atoi(path[len(path)-1])
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	student, err := api.usecase.GetStudentByID(studentID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(&m.StudentByIDResponse{Student: *student})
}

// GetStudentsFromClass godoc
// @Summary Get students from class
// @Description Get students from class by class id
// @ID getStudentsFromClass
// @Accept  json
// @Produce  json
// @Tags Student
// @Param classID path string true "Class id"
// @Success 200 {object} m.StudentListFromClass
// @Failure 400 {object} m.Error "bad request - Problem with the request"
// @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} m.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} m.Error "internal server error - Request is valid but operation failed at server side"
// @Router /classes/{classID}/students [get]
func (api *Handler) GetStudentsFromClass(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	chatId, err := strconv.Atoi(path[len(path)-2])
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	students, err := api.usecase.GetStudentsFromClass(chatId)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(&m.StudentListFromClass{Students: students})
}
