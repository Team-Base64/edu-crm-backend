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

// GetSolutionsFromClass godoc
// @Summary Get solutions from class
// @Description Get solutions from class by class id
// @ID getSolutionsFromClass
// @Accept  json
// @Produce  json
// @Tags Solution
// @Param classID path string true "Class id"
// @Success 200 {object} model.SolutionListFromClass
// @Failure 400 {object} model.Error "bad request - Problem with the request"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} model.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /classes/{classID}/solutions [get]
func (api *Handler) GetSolutionsFromClass(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	classID, err := strconv.Atoi(path[len(path)-2])
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	sols, err := api.usecase.GetSolutionsByClassID(classID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(sols)
}

// GetSolutionsForHomework godoc
// @Summary Get solutions for homework
// @Description Get solutions for homework by homework id
// @ID getSolutionsForHomework
// @Accept  json
// @Produce  json
// @Tags Solution
// @Param homeworkID path string true "Homework id"
// @Success 200 {object} model.SolutionListForHw
// @Failure 400 {object} model.Error "bad request - Problem with the request"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} model.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /homeworks/{homeworkID}/solutions [get]
func (api *Handler) GetSolutionsForHomework(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	hwID, err := strconv.Atoi(path[len(path)-2])
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	sols, err := api.usecase.GetSolutionsByHomeworkID(hwID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(sols)
}

// GetSolution godoc
// @Summary Get solution
// @Description Get solution by id
// @ID getSolution
// @Accept  json
// @Produce  json
// @Tags Solution
// @Param solID path string true "Solution id"
// @Success 200 {object} model.SolutionByIDResponse
// @Failure 400 {object} model.Error "bad request - Problem with the request"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 404 {object} model.Error "not found - Requested entity is not found in database"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /solutions/{solID} [get]
func (api *Handler) GetSolution(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")
	solID, err := strconv.Atoi(path[len(path)-1])
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	sol, err := api.usecase.GetSolutionByID(solID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(&model.SolutionByIDResponse{Solution: *sol})
}
