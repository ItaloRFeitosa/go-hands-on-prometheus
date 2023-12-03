package errs

import "net/http"

var statusCodeMap = map[Type]int{
	InternalType:         http.StatusInternalServerError,
	NotFoundType:         http.StatusNotFound,
	InvalidOperationType: http.StatusUnprocessableEntity,
	InvalidParamsType:    http.StatusBadRequest,
	ConflictType:         http.StatusConflict,
}

func HttpStatusCode(err error) int {
	asError, ok := AsError(err)
	if !ok {
		return http.StatusInternalServerError
	}

	code, ok := statusCodeMap[asError.Type]

	if !ok {
		return http.StatusInternalServerError
	}

	return code
}
