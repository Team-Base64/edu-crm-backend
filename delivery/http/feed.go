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

// GetClassFeed godoc
// @Summary Get class feed
// @Description Get posts from class by class id
// @ID getClassFeed
// @Accept  json
// @Produce  json
// @Tags Feed
// @Param classID path string true "Class id"
// @Success 200 {object} m.Feed
// @Failure 400 {object} m.Error "bad request - Problem with the request"
// @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} m.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} m.Error "internal server error - Request is valid but operation failed at server side"
// @Router /classes/{classID}/feed [get]
func (api *Handler) GetClassFeed(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	chatId, err := strconv.Atoi(path[len(path)-2])
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	posts, err := api.usecase.GetClassPosts(chatId)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(&m.Feed{Posts: posts})
}

// CreatePost godoc
// @Summary Create post
// @Description Create post
// @ID createPost
// @Accept  json
// @Produce  json
// @Tags Feed
// @Param classID path string true "Class id"
// @Param post body m.PostCreate true "Post for creating"
// @Success 200 {object} m.PostResponse
// @Failure 400 {object} m.Error "bad request - Problem with the request"
// @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} m.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} m.Error "internal server error - Request is valid but operation failed at server side"
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
	var newPost m.PostCreate
	if err := decoder.Decode(&newPost); err != nil {
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}
	sanitizer := bluemonday.UGCPolicy()
	newPost.Text = sanitizer.Sanitize(newPost.Text)
	post, err := api.usecase.CreatePost(classID, &newPost)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(&m.PostResponse{Post: *post})
}
