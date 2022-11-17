package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/ong-gtp/go-chat/pkg/domain/responses"
	"github.com/ong-gtp/go-chat/pkg/errors"
)

func ParseBody(r *http.Request, x interface{}) error {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return err
		}
	}
	return nil
}

func Ok(res []byte, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

type emptyOk struct {
	Message string
}

func OkEmpty(message string, w http.ResponseWriter) {
	m := emptyOk{message}
	res, err := json.Marshal(m)
	errors.ErrorCheck(err)
	Ok(res, w)
}

// func DerefString(s *string) string {
// 	if s != nil {
// 		return *s
// 	}
// 	return ""
// }

func ErrResponse(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	errCode := codeFrom(err)
	w.WriteHeader(errCode)
	res := responses.ErrorResponse{Message: err.Error(), Status: false, Code: errCode}
	data, err := json.Marshal(res)
	errors.ErrorCheck(err)
	w.Write(data)
}

// codeFrom returns the http status code from service errors
func codeFrom(err error) int {
	switch err {
	case errors.ErrInvalidCredentials:
		return http.StatusBadRequest
	case errors.ErrDuplicateEmail:
		return http.StatusBadRequest
	case errors.ErrInRequestMarshaling:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func Revive() {
	if err := recover(); err != nil {
		if os.Getenv("LOG_PANIC_TRACE") == "true" {
			log.Println(
				"level:", "error",
				"err: ", err,
				"trace", string(debug.Stack()),
			)
		} else {
			log.Println(
				"level", "error",
				"err", err,
			)
		}
	}
}
