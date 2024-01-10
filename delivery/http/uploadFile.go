package delivery

import (
	"encoding/json"
	"log"
	"net/http"

	e "main/domain/errors"
	m "main/domain/model"
)

// UploadAttach godoc
// @Summary Upload attach
// @Description Upload attach
// @ID uploadAttach
// @Accept  multipart/form-data
// @Produce  json
// @Param file formData file true "attach"
// @Param type query string true "type: homework or solution or chat"
// @Success 200 {object} m.Response "ok"
// @Failure 401 {object} m.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} m.Error "internal Server Error - Request is valid but operation failed at server side"
// @Router /attach [post]
func (api *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, e.ErrBadRequest400)
		return
	}
	defer file.Close()

	fileName, err := api.usecase.SaveAttach(&m.Attach{
		Dest: r.URL.Query().Get("type"),
		File: file,
	})
	if err != nil {
		log.Println(e.StacktraceError(err))
		returnErrorJSON(w, err)
		return
	}

	json.NewEncoder(w).Encode(&m.UploadAttachResponse{File: fileName})
}
