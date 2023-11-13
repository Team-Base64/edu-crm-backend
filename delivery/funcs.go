package delivery

import (
	"encoding/json"
	"net/http"

	"main/domain/model"

	e "main/domain/errors"
)

var mockTeacherID = 1

func ReturnErrorJSON(w http.ResponseWriter, err error) {
	errCode, errText := e.CheckError(err)
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(&model.Error{Error: errText})
}
