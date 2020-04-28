package responses

import (
	"encoding/json"
	"net/http"
)

const (
	CONTENT_TYPE_JSON = "application/json"
)

type jsonFormatter struct{}

func (response *jsonFormatter) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", CONTENT_TYPE_JSON)
	w.WriteHeader(code)

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(data); err != nil {
		encoder.Encode(serializeErrorJson(err))
	}
}

func (response *jsonFormatter) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	response.respond(w, r, code, serializeErrorJson(err))
}

func newJsonFormatter() *jsonFormatter {
	return &jsonFormatter{}
}

func serializeErrorJson(err error) map[string]string {
	return map[string]string{
		"error": err.Error(),
	}
}
