package delivery

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	e "main/domain/errors"
	"main/domain/model"
	"net/http"
	"os"

	"github.com/google/uuid"
)

// UploadAttach godoc
// @Summary Upload attach
// @Description Upload attach
// @ID uploadAttach
// @Accept  multipart/form-data
// @Produce  json
// @Param file formData file true "attach"
// @Param type query string true "type: homework or solution or chat"
// @Success 200 {object} model.Response "ok"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal Server Error - Request is valid but operation failed at server side"
// @Router /attach [post]
func (api *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	typeS := r.URL.Query().Get("type")

	filePath := ""
	switch typeS {
	case "homework":
		filePath = api.filestoragePath + homeworkFilesPath
	case "solution":
		filePath = api.filestoragePath + solutionFilesPath
	case "chat":
		filePath = api.filestoragePath + chatFilesPath
	default:
		log.Println(e.StacktraceError(errors.New("error wrong type query param")))
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}
	defer file.Close()

	fileHeader := make([]byte, 512)

	// Copy the headers into the FileHeader buffer
	if _, err := file.Read(fileHeader); err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	// set position back to start.
	if _, err := file.Seek(0, 0); err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}
	log.Println(http.DetectContentType(fileHeader))
	fileExt := ""
	switch http.DetectContentType(fileHeader) {
	case "image/jpeg":
		fileExt = ".jpg"
	case "image/png":
		fileExt = ".png"
	case "application/pdf":
		fileExt = ".pdf"
	case "application/vnd.rar":
		fileExt = ".rar"
	case "application/x-rar-compressed":
		fileExt = ".rar"
	case "application/zip":
		fileExt = ".zip"
	case "application/x-zip-compressed":
		fileExt = ".zip"
	default:
		log.Println("error not allowed file extension")
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	attachNum := uuid.New().String()

	fileName := filePath + "/" + attachNum + fileExt
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrServerError500)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrServerError500)
		return
	}

	json.NewEncoder(w).Encode(&model.UploadAttachResponse{File: api.urlDomain + fileName})
}
