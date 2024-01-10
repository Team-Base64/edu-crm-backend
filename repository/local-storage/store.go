package localstorage

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	e "main/domain/errors"
	m "main/domain/model"
	rep "main/repository"

	"github.com/google/uuid"
)

type LocalStore struct {
	chatFilesPath     string
	homeworkFilesPath string
	solutionFilesPath string
	filestoragePath   string
}

func NewLocalStore(cfp string, hfp string, sfp string, fsp string) rep.FileStoreInterface {
	for _, path := range []string{cfp, hfp, sfp} {
		if err := os.MkdirAll(fsp+path, os.ModePerm); err != nil {
			log.Fatalln(e.StacktraceError(err))
		}
	}

	return &LocalStore{
		chatFilesPath:     cfp,
		homeworkFilesPath: hfp,
		solutionFilesPath: sfp,
		filestoragePath:   fsp,
	}
}

func (s *LocalStore) UploadFile(file *m.Attach) (string, error) {
	filePath := ""
	switch file.Dest {
	case "homework":
		filePath = s.filestoragePath + s.homeworkFilesPath
	case "solution":
		filePath = s.filestoragePath + s.solutionFilesPath
	case "chat":
		filePath = s.filestoragePath + s.chatFilesPath
	default:
		return "", e.StacktraceError(errors.New("error wrong destination"), e.ErrBadRequest400)
	}

	fileHeader := make([]byte, 512)

	// Copy the headers into the FileHeader buffer
	if _, err := file.File.Read(fileHeader); err != nil {
		return "", e.StacktraceError(err, e.ErrBadRequest400)
	}

	// set position back to start.
	if _, err := file.File.Seek(0, 0); err != nil {
		return "", e.StacktraceError(err, e.ErrBadRequest400)
	}

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
		return "", e.StacktraceError(errors.New("error not allowed file extension"), e.ErrBadRequest400)
	}

	fileName := filePath + "/" + uuid.New().String() + fileExt
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", e.StacktraceError(err)
	}
	defer f.Close()

	_, err = io.Copy(f, file.File)
	if err != nil {
		return "", e.StacktraceError(err)
	}

	return fileName, nil
}
