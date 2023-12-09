package delivery

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	e "main/domain/errors"
	"main/domain/model"
	"net/http"
	"time"

	"github.com/microcosm-cc/bluemonday"
)

// SignUp godoc
// @Summary Sign Up and returns the authentication  cookie
// @Description Sign Up user
// @ID signup
// @Accept  json
// @Produce  json
// @Tags Teacher
// @Param teacher body model.TeacherSignUp true "Teacher params"
// @Success 201 {object} model.Response "OK"
// @Failure 400 {object} model.Error "bad request - Problem with the request"
// @Failure 409 {object} model.Error "conflict - UserDB already exists"
// @Failure 500 {object} model.Error "internal Server Error - Request is valid but operation failed at server side"
// @Failure 503 {object} model.Error "service unavailable"
// @Router /register [post]
func (api *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req model.TeacherSignUp
	if err := decoder.Decode(&req); err != nil {
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	sanitizer := bluemonday.UGCPolicy()
	req.Login = sanitizer.Sanitize(req.Login)
	req.Name = sanitizer.Sanitize(req.Name)

	user, err := api.usecase.GetTeacherProfileByLogin(req.Login)
	if user != nil && user.Login != "" {
		returnErrorJSON(w, e.ErrConflict409)
		return
	}

	err = api.usecase.SignUpTeacher(&req)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	sess, err := api.usecase.CreateSession(req.Login)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrServerError500)
		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sess.ID,
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(201)

	json.NewEncoder(w).Encode(&model.Response{})
}

// LogIn godoc
// @Summary Logs in and returns the authentication  cookie
// @Description Log in user
// @ID login
// @Accept  json
// @Produce  json
// @Tags Teacher
// @Param teacher body model.TeacherLogin true "Teacher params"
// @Success 201 {object} model.Response "OK"
// @Failure 400 {object} model.Error "bad request - Problem with the request"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal Server Error - Request is valid but operation failed at server side"
// @Router /login [post]
func (api *Handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var req model.TeacherLogin
	if err := decoder.Decode(&req); err != nil {
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	user, err := api.usecase.GetTeacherProfileByLogin(req.Login)
	if errors.Is(err, sql.ErrNoRows) {
		log.Println(e.StacktraceError(err, errors.New("no user: "+req.Login)))
		returnErrorJSON(w, e.ErrUnauthorized401)
		return
	}
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrServerError500)
		return
	}
	if user.Password != req.Password {
		log.Println(e.StacktraceError(errors.New("wrong password")))
		returnErrorJSON(w, e.ErrUnauthorized401)
		return
	}

	sess, err := api.usecase.CreateSession(user.Login)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrServerError500)
		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sess.ID,
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(201)

	json.NewEncoder(w).Encode(&model.Response{})
}

// LogOut godoc
// @Summary Logs out user
// @Description Logs out user
// @ID logout
// @Accept  json
// @Produce  json
// @Tags Teacher
// @Success 200 {object} model.Response "OK"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal Server Error - Request is valid but operation failed at server side"
// @Router /logout [delete]
func (api *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	session, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}
	// _, err = api.CheckSession(session.Value)
	// if err != nil {
	// 	log.Println(e.StacktraceError(err))
	// 	returnErrorJSON(w, e.ErrUnauthorized401)
	// 	return
	// }
	err = api.usecase.DeleteSession(session.Value)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrServerError500)
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	json.NewEncoder(w).Encode(&model.Response{})
}

// Auth godoc
// @Summary Check user auth
// @Description Check user auth
// @ID auth
// @Accept  json
// @Produce  json
// @Tags Teacher
// @Success 200 {object} model.Response "OK"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal Server Error - Request is valid but operation failed at server side"
// @Router /auth [get]
func (api *Handler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	_, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrUnauthorized401)
		return
	}
	json.NewEncoder(w).Encode(&model.Response{})
}
