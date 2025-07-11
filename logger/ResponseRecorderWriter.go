package logger

import (
	"net/http"
)

type ResponseRecorderWriter struct {
	http.ResponseWriter
	Status int
}

func (rr *ResponseRecorderWriter) WriteHeader(status int) {
	rr.Status = status
	rr.ResponseWriter.WriteHeader(status)
}

func (rr *ResponseRecorderWriter) Flush() {
	rr.ResponseWriter.(http.Flusher).Flush()
}
