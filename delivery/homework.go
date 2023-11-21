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

// GetHomeworksFromClass godoc
// @Summary Get class homeworks
// @Description Get homeworks from class by class id
// @ID getHomeworksFromClass
// @Accept  json
// @Produce  json
// @Tags Homework
// @Param classID path string true "Class id"
// @Success 200 {object} model.HomeworkList
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
// @Param hwID path string true "Homework id"
// @Success 200 {object} model.HomeworkByIDResponse
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
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	hw, err := api.usecase.GetHomeworkByID(hwID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(&model.HomeworkByIDResponse{Homework: *hw})
}

// CreateHomework godoc
// @Summary Create homework
// @Description Create homework
// @ID createHomework
// @Accept  json
// @Produce  json
// @Tags Homework
// @Param post body model.HomeworkCreate true "Homework for creating"
// @Success 200 {object} model.HomeworkResponse
// @Failure 400 {object} model.Error "bad request - Problem with the request"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} model.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /homeworks [post]
func (api *Handler) CreateHomework(w http.ResponseWriter, r *http.Request) {
	teacherProfile := r.Context().Value(KeyUserdata{"userdata"}).(*model.TeacherDB)
	decoder := json.NewDecoder(r.Body)
	var newHw model.HomeworkCreate
	if err := decoder.Decode(&newHw); err != nil {
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	hw, err := api.usecase.CreateHomework(teacherProfile.ID, &newHw)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(&model.HomeworkResponse{Homework: *hw})
}
