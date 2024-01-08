package delivery

import (
	"encoding/json"
	"log"
	e "main/domain/errors"
	"main/domain/model"
	"net/http"
)

// // CreateTeacher godoc
// // @Summary Create teacher
// // @Description Create teacher
// // @ID createTeacher
// // @Accept  json
// // @Produce  json
// // @Tags Teacher
// // @Param user body model.TeacherSignUp true "Teacher params"
// // @Success 200 {object} model.Response "OK"
// // @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// // @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// // @Router /register2 [post]
// func (api *Handler) CreateTeacher(w http.ResponseWriter, r *http.Request) {
// 	decoder := json.NewDecoder(r.Body)
// 	var req model.TeacherSignUp
// 	if err := decoder.Decode(&req); err != nil {
// 		returnErrorJSON(w, e.ErrBadRequest400)
// 		return
// 	}
// 	sanitizer := bluemonday.UGCPolicy()
// 	req.Login = sanitizer.Sanitize(req.Login)
// 	req.Name = sanitizer.Sanitize(req.Name)
// 	if err := api.usecase.CreateTeacher(&req); err != nil {
// 		log.Println(e.StacktraceError(err))
// 		returnErrorJSON(w, err)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(&model.Response{})
// }

// GetTeacher godoc
// @Summary Get teacher's info
// @Description gets teacher's info
// @ID getTeacher
// @Accept  json
// @Produce  json
// @Tags Teacher
// @Success 200 {object} model.TeacherProfileResponse
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /profile [get]
func (api *Handler) GetTeacherProfile(w http.ResponseWriter, r *http.Request) {
	teacherProfile := r.Context().Value(KeyUserdata{"userdata"}).(*model.TeacherDB)
	teacher, err := api.usecase.GetTeacherProfile(teacherProfile.ID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(&model.TeacherProfileResponse{Teacher: *teacher})
}
