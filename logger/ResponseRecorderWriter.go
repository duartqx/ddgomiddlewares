package logger

import (
	"net/http"
	"strings"
)

type ResponseRecorderWriter struct {
	http.ResponseWriter
	Status int
	Body   []string
}

func (rr *ResponseRecorderWriter) WriteHeader(status int) {
	rr.Status = status
	rr.ResponseWriter.WriteHeader(status)
}

func (rr *ResponseRecorderWriter) Write(msg []byte) (int, error) {
	rr.Body = append(rr.Body, string(msg))
	return rr.ResponseWriter.Write(msg)
}

func (rr ResponseRecorderWriter) BodyString() string {
	return strings.Join(rr.Body, "; ")
}
