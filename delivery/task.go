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
	teacherProfile := r.Context().Value(KeyUserdata{"userdata"}).(*model.TeacherDB)
	tasks, err := api.usecase.GetTasksByTeacher(teacherProfile.ID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}
