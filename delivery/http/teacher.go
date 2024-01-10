package delivery

import (
	"encoding/json"
	"log"
	"net/http"

	e "main/domain/errors"
	m "main/domain/model"
)

// // CreateTeacher godoc
// // @Summary Create teacher
// // @Description Create teacher
// // @ID createTeacher
// // @Accept  json
// // @Produce  json
// // @Tags Teacher
// // @Param user body m.TeacherSignUp true "Teacher params"
// // @Success 200 {object} m.Response "OK"
// // @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// // @Failure 500 {object} m.Error "internal server error - Request is valid but operation failed at server side"
// // @Router /register2 [post]
// func (api *Handler) CreateTeacher(w http.ResponseWriter, r *http.Request) {
// 	decoder := json.NewDecoder(r.Body)
// 	var req m.TeacherSignUp
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

// 	json.NewEncoder(w).Encode(&m.Response{})
// }

// GetTeacher godoc
// @Summary Get teacher's info
// @Description gets teacher's info
// @ID getTeacher
// @Accept  json
// @Produce  json
// @Tags Teacher
// @Success 200 {object} m.TeacherProfileResponse
// @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} m.Error "internal server error - Request is valid but operation failed at server side"
// @Router /profile [get]
func (api *Handler) GetTeacherProfile(w http.ResponseWriter, r *http.Request) {
	teacherProfile := r.Context().Value(KeyUserdata{"userdata"}).(*m.TeacherDB)
	teacher, err := api.usecase.GetTeacherProfile(teacherProfile.ID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(&m.TeacherProfileResponse{Teacher: *teacher})
}
