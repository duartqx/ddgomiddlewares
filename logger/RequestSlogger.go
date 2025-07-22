package logger

import (
	"net/http"
	"time"

	"github.com/duartqx/ddgomiddlewares/logger/colors"
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
			"Method", rl.getColored(rl.method),
			"Status", rl.getColored(rl.status),
			"Since", rl.since.String(),
			"Path", rl.path,
			"Error", rl.getColored(rl.result),
		}
	}
	return []any{
		"Method", rl.getColored(rl.method),
		"Status", rl.getColored(rl.status),
		"Since", rl.since.String(),
		"Path", rl.path,
	}
}

func (rl RequestSLogger) getColored(value any) string {
	return colors.GetStatusColor(rl.status, value)
}
