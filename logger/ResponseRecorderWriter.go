package logger

import (
	"net/http"
)

type ResponseRecorderWriter struct {
	http.ResponseWriter
	Status       int
	WroteHeaders bool
	Written      *struct {
		Res int
		Err error
	}
	Result string
}

func (rr *ResponseRecorderWriter) Write(b []byte) (int, error) {
	if rr.Header().Get("Content-Type") != "application/json" {
		return rr.ResponseWriter.Write(b)
	}

	if rr.Written != nil {
		return rr.Written.Res, rr.Written.Err
	}

	rr.Result = string(b)

	rr.Written = &struct {
		Res int
		Err error
	}{}

	rr.WriteHeader(http.StatusOK)

	rr.Written.Res, rr.Written.Err = rr.ResponseWriter.Write(b)

	return rr.Written.Res, rr.Written.Err
}

func (rr *ResponseRecorderWriter) WriteHeader(status int) {
	if rr.WroteHeaders {
		return
	}

	rr.Status = status
	rr.ResponseWriter.WriteHeader(status)
	rr.WroteHeaders = true
}

func (rr *ResponseRecorderWriter) Flush() {
	rr.ResponseWriter.(http.Flusher).Flush()
}
