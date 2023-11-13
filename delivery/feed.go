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

// GetClassFeed godoc
// @Summary Get class feed
// @Description Get posts from class by class id
// @ID getClassFeed
// @Accept  json
// @Produce  json
// @Param classID path string true "Class id"
// @Success 200 {object} model.Feed
// @Failure 400 {object} model.Error "bad request - Problem with the request"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} model.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /classes/{classID}/feed [get]
func (api *Handler) GetClassFeed(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	chatId, err := strconv.Atoi(path[len(path)-2])
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	feed, err := api.usecase.GetClassFeed(chatId)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(feed)
}

// CreatePost godoc
// @Summary Create post
// @Description Create post
// @ID createPost
// @Accept  json
// @Produce  json
// @Param classID path string true "Class id"
// @Param post body model.PostCreate true "Post for creating"
// @Success 200 {object} model.PostResponse
// @Failure 400 {object} model.Error "bad request - Problem with the request"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} model.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /classes/{classID}/feed [post]
func (api *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	classID, err := strconv.Atoi(path[len(path)-2])
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var newPost model.PostCreate
	if err := decoder.Decode(&newPost); err != nil {
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	post, err := api.usecase.CreatePost(classID, &newPost)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(&model.PostResponse{Post: *post})
}
