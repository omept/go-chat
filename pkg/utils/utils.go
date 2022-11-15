package utils

import (
	"encoding/json"
	"io"
	"net/http"

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
