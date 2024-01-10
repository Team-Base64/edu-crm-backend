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

// GetTeacherTasks godoc
// @Summary Get teacher's tasks
// @Description gets teacher's tasks
// @ID getTeacherTasks
// @Accept  json
// @Produce  json
// @Tags Tasks
// @Success 200 {object} m.TaskListByTeacherID
// @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} m.Error "internal server error - Request is valid but operation failed at server side"
// @Router /tasks [get]
func (api *Handler) GetTeacherTasks(w http.ResponseWriter, r *http.Request) {
	teacherProfile := r.Context().Value(KeyUserdata{"userdata"}).(*m.TeacherDB)
	tasks, err := api.usecase.GetTasksByTeacherID(teacherProfile.ID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(m.TaskListByTeacherID{Tasks: tasks})
}

// CreateTasks godoc
// @Summary Create task by teacher
// @Description create task by teacher
// @ID createTasks
// @Accept  json
// @Produce  json
// @Tags Tasks
// @Param post body m.TaskCreate true "Task for creating"
// @Success 200 {object} m.TaskCreateResponse
// @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} m.Error "internal server error - Request is valid but operation failed at server side"
// @Router /tasks [post]
func (api *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req m.TaskCreate
	if err := decoder.Decode(&req); err != nil {
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}
	sanitizer := bluemonday.UGCPolicy()
	req.Description = sanitizer.Sanitize(req.Description)

	teacherProfile := r.Context().Value(KeyUserdata{"userdata"}).(*m.TeacherDB)
	task, err := api.usecase.CreateTask(teacherProfile.ID, &req)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(task)
}

// GetTaskByID godoc
// @Summary Get task by ID
// @Description get task by its ID
// @ID getTaskByID
// @Accept  json
// @Produce  json
// @Tags Tasks
// @Param taskID path string true "Task id"
// @Success 200 {object} m.TaskByIDResponse
// @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} m.Error "internal server error - Request is valid but operation failed at server side"
// @Router /tasks/{taskID} [get]
func (api *Handler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	taskID, err := strconv.Atoi(path[len(path)-1])
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	task, err := api.usecase.GetTaskByID(taskID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(&m.TaskByIDResponse{Task: *task})
}
