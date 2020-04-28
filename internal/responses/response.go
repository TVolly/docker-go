package responses

import (
	"encoding/json"
	"net/http"
)

const (
	CONTENT_TYPE_JSON = "application/json"
)

var (
	onFireError FireErrorFunc
)

type FireErrorFunc func(error, int)

func Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", CONTENT_TYPE_JSON)
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	message, isSecure := getSecureMessage(err)

	if !isSecure {
		fireError(err, code)
	}

	Respond(w, r, code, map[string]string{
		"error": message,
	})
}

func SetFireError(f FireErrorFunc) {
	onFireError = f
}

func getSecureMessage(err error) (string, bool) {
	switch err {
	case ErrPageNorFound:
		return err.Error(), true
	case ErrInvalidData:
		return err.Error(), true
	}
	return "System error", false
}

func fireError(err error, code int) {
	if onFireError != nil {
		onFireError(err, code)
	}
}
