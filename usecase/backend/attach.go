package backend

import (
	e "main/domain/errors"
	m "main/domain/model"
)

func (uc *BackendUsecase) SaveAttach(file *m.Attach) (string, error) {
	fileName, err := uc.fileStore.UploadFile(file)
	if err != nil {
		return "", e.StacktraceError(err)
	}
	return uc.urlDomain + fileName, nil
}
