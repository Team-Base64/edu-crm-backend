package delivery

import (
	"encoding/json"
	"log"
	e "main/domain/errors"
	"main/domain/model"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (api *Handler) CreateSession(in *model.Session) (string, error) {
	log.Println("call Create session", in)
	newUUID := uuid.New()
	id := newUUID.String()
	api.mu.Lock()
	api.sessions[id] = in
	api.mu.Unlock()

	return id, nil
}

func (api *Handler) CheckSession(in string) (*model.Session, error) {
	log.Println("call Check Session", in)
	api.mu.RLock()
	defer api.mu.RUnlock()
	if sess, ok := api.sessions[in]; ok {
		return sess, nil
	}
	return nil, e.ErrUnauthorized401
}

func (api *Handler) DeleteSession(in string) error {
	log.Println("call Delete Session", in)
	api.mu.Lock()
	defer api.mu.Unlock()
	delete(api.sessions, in)
	return nil
}

// LogIn godoc
// @Summary Logs in and returns the authentication  cookie
// @Description Log in user
// @ID login
// @Accept  json
// @Produce  json
// @Tags User
// @Param user body model.TeacherLogin true "Teacher params"
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

	userID, err := api.usecase.CheckLogin(req)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	sess, err := api.CreateSession(userID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrServerError500)
		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    sess,
		Expires:  time.Now().Add(10 * time.Hour),
		HttpOnly: true,
		Secure:   true,
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
// @Tags User
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
	err = api.DeleteSession(session.Value)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrServerError500)
		return
	}
	session.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, session)
	json.NewEncoder(w).Encode(&model.Response{})
}
