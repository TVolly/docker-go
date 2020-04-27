package responses

import (
	"encoding/json"
	"net/http"
)

type jsonFormatter struct{}

func (response *jsonFormatter) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (response *jsonFormatter) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	response.respond(w, r, code, map[string]string{"error": err.Error()})
}

func newJsonFormatter() *jsonFormatter {
	return &jsonFormatter{}
}
