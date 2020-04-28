package middlewares

import (
	"net/http"
	"reflect"
	"time"

	"github.com/TVolly/goapi-addresses/internal/responses"
	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
)

type logRequest struct {
	logger *logrus.Logger
}

func (m *logRequest) handler(next http.Handler) http.Handler {
	responses.SetFireError(func(err error, code int) {
		m.logger.Errorf("%s: %s", reflect.TypeOf(err), err.Error())
	})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := m.logger.WithFields(logrus.Fields{
			"ip": r.RemoteAddr,
		})

		start := time.Now()
		logger.Tracef("[%s] %s", r.Method, r.RequestURI)

		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Tracef(
			"%d %s | %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}

func LogRequests(l *logrus.Logger) mux.MiddlewareFunc {
	m := &logRequest{logger: l}

	return m.handler
}
