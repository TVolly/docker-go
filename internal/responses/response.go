package responses

import "net/http"

type formatter interface {
	respond(w http.ResponseWriter, r *http.Request, code int, data interface{})
	error(w http.ResponseWriter, r *http.Request, code int, err error)
}

var (
	jsonFormat = newJsonFormatter()
)

func Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	jsonFormat.respond(w, r, code, data)
}

func Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	jsonFormat.error(w, r, code, err)
}
