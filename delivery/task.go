package delivery

import (
	"encoding/json"
	"log"
	e "main/domain/errors"
	"main/domain/model"
	"net/http"
)

// GetTeacherTasks godoc
// @Summary Get teacher's tasks
// @Description gets teacher's tasks
// @ID getTeacherTasks
// @Accept  json
// @Produce  json
// @Tags Tasks
// @Success 200 {object} model.TaskListByTeacherID
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /tasks [get]
func (api *Handler) GetTeacherTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := api.usecase.GetTasksByTeacherID(mockTeacherID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(model.TaskListByTeacherID{Tasks: tasks})
}

// GetTeacherTasks godoc
// @Summary Get teacher's tasks
// @Description gets teacher's tasks
// @ID getTeacherTasks
// @Accept  json
// @Produce  json
// @Tags Tasks
// @Success 200 {object} model.TaskListByTeacherID
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /tasks [post]
func (api *Handler) CreateTasks(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req model.TaskCreate
	if err := decoder.Decode(&req); err != nil {
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	task, err := api.usecase.CreateTask(mockTeacherID, &req)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(task)
}
