package middlewares

import (
	"fmt"
	"log/slog"
	"net/http"
)

type statusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &statusRecorder{
			ResponseWriter: w,
			Status:         200,
		}

		next.ServeHTTP(rec, r)

		logMsg := "Request"
		if rec.Status == http.StatusInternalServerError || rec.Status == http.StatusBadRequest {
			logMsg = "Request RESPONSIBLE FOR THE LAST ERROR"
		}
		slog.Info(
			logMsg,
			"host", r.RemoteAddr,
			"protocol", r.Proto,
			"method", r.Method,
			"url", r.URL.RequestURI(),
			"status", fmt.Sprintf("%d (%s)", rec.Status, http.StatusText(rec.Status)))
	})
}
