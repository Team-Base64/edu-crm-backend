package BaseErrors

import (
	"errors"
	"runtime"
	"strconv"
)

var ErrBadRequest400 = errors.New("bad request - Problem with the request")
var ErrUnauthorized401 = errors.New("unauthorized - Access token is missing or invalid")
var ErrForbidden403 = errors.New("forbidden")
var ErrNotFound404 = errors.New("not found - Requested entity is not found in database")
var ErrConflict409 = errors.New("conflict - UserDB already exists")
var ErrServerError500 = errors.New("internal server error - Request is valid but operation failed at server side")

func CheckBaseError(err error) (int, error) {
	switch err {
	case ErrBadRequest400:
		return 400, ErrBadRequest400
	case ErrUnauthorized401:
		return 401, ErrUnauthorized401
	case ErrForbidden403:
		return 403, ErrForbidden403
	case ErrNotFound404:
		return 404, ErrNotFound404
	case ErrConflict409:
		return 409, ErrConflict409
	case ErrServerError500:
		return 500, ErrServerError500
	}

	return -1, err
}

func StacktraceError(err error) error {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return err
	}
	return errors.Join(
		err,
		errors.New("				at "+file+":"+strconv.Itoa(line)),
	)
}
