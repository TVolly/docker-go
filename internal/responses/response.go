package responses

import "net/http"

type formatter interface {
	respond(w http.ResponseWriter, r *http.Request, code int, data interface{})
	error(w http.ResponseWriter, r *http.Request, code int, err error)
}

var (
	jsonFormat = newJsonFormatter()
)

func getFormatter(r *http.Request) formatter {
	return jsonFormat
}

func Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	getFormatter(r).respond(w, r, code, data)
}

func Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	getFormatter(r).error(w, r, code, err)
}
