package delivery

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	e "main/domain/errors"
	m "main/domain/model"

	"github.com/microcosm-cc/bluemonday"
)

// GetHomeworksFromClass godoc
// @Summary Get class homeworks
// @Description Get homeworks from class by class id
// @ID getHomeworksFromClass
// @Accept  json
// @Produce  json
// @Tags Homework
// @Param classID path string true "Class id"
// @Success 200 {object} m.HomeworkList
// @Failure 400 {object} m.Error "bad request - Problem with the request"
// @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} m.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} m.Error "internal server error - Request is valid but operation failed at server side"
// @Router /classes/{classID}/homeworks [get]
func (api *Handler) GetHomeworksFromClass(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	chatId, err := strconv.Atoi(path[len(path)-2])
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	hws, err := api.usecase.GetHomeworksByClassID(chatId)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
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
// @Tags Homework
// @Param homeworkID path string true "Homework id"
// @Success 200 {object} m.HomeworkByIDResponse
// @Failure 400 {object} m.Error "bad request - Problem with the request"
// @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} m.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} m.Error "internal server error - Request is valid but operation failed at server side"
// @Router /homeworks/{homeworkID} [get]
func (api *Handler) GetHomework(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	hwID, err := strconv.Atoi(path[len(path)-1])
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	hw, err := api.usecase.GetHomeworkByID(hwID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(&m.HomeworkByIDResponse{Homework: *hw})
}

// CreateHomework godoc
// @Summary Create homework
// @Description Create homework
// @ID createHomework
// @Accept  json
// @Produce  json
// @Tags Homework
// @Param post body m.HomeworkCreate true "Homework for creating"
// @Success 200 {object} m.HomeworkResponse
// @Failure 400 {object} m.Error "bad request - Problem with the request"
// @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} m.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} m.Error "internal server error - Request is valid but operation failed at server side"
// @Router /homeworks [post]
func (api *Handler) CreateHomework(w http.ResponseWriter, r *http.Request) {
	teacherProfile := r.Context().Value(KeyUserdata{"userdata"}).(*m.TeacherDB)
	decoder := json.NewDecoder(r.Body)
	var newHw m.HomeworkCreate
	if err := decoder.Decode(&newHw); err != nil {
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}
	sanitizer := bluemonday.UGCPolicy()
	newHw.Title = sanitizer.Sanitize(newHw.Title)
	newHw.Description = sanitizer.Sanitize(newHw.Description)
	hw, err := api.usecase.CreateHomework(teacherProfile.ID, &newHw)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(&m.HomeworkResponse{Homework: *hw})
}
