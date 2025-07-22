package logger

import (
	"net/http"
	"time"
)

type RequestSLogger struct {
	method string
	result any
	status int
	since  time.Duration
	path   string
}

func NewRequestSLogger() *RequestSLogger {
	return &RequestSLogger{}
}

func (rl RequestSLogger) WithMethod(method string) RequestSLogger {
	rl.method = method
	return rl
}

func (rl RequestSLogger) WithResult(result any) RequestSLogger {
	rl.result = result
	return rl
}

func (rl RequestSLogger) WithStatus(status int) RequestSLogger {
	rl.status = status
	return rl
}

func (rl RequestSLogger) WithSince(since time.Duration) RequestSLogger {
	rl.since = since
	return rl
}

func (rl RequestSLogger) WithPath(path string) RequestSLogger {
	rl.path = path
	return rl
}

func (rl RequestSLogger) Slog() []any {
	if rl.status >= http.StatusBadRequest {
		return []any{
			"method", rl.method,
			"status", rl.status,
			"since", rl.since.String(),
			"path", rl.path,
			"error", rl.result,
		}
	}
	return []any{
		"method", rl.method,
		"status", rl.status,
		"since", rl.since.String(),
		"path", rl.path,
	}
}
