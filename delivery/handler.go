package delivery

import (
	"encoding/json"
	"log"
	e "main/domain/errors"
	"main/domain/model"
	uc "main/usecase"
	"net/http"
	"os"
)

// @title TCRA API
// @version 1.0
// @description TCRA back server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath  /api

var chatFilesPath = "/chat"
var homeworkFilesPath = "/homework"
var solutionFilesPath = "/solution"

type Handler struct {
	usecase     uc.UsecaseInterface
	filestorage string
}

func NewHandler(uc uc.UsecaseInterface, fs string) *Handler {
	for _, path := range []string{chatFilesPath, homeworkFilesPath, solutionFilesPath} {
		if err := os.MkdirAll(fs+path, os.ModePerm); err != nil {
			log.Fatalln(e.StacktraceError(err))
		}
	}

	return &Handler{
		usecase: uc,

		filestorage: fs,
	}
}

var mockTeacherID = 1

func returnErrorJSON(w http.ResponseWriter, err error) {
	errCode, errText := e.CheckError(err)
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(&model.Error{Error: errText})
}
