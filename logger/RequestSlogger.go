package logger

import (
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/google/uuid"
)

type RequestSLogger struct {
	id      uuid.UUID
	method  string
	result  any
	status  int
	since   time.Duration
	host    string
	path    string
	service string
}

func NewRequestSLogger(id uuid.UUID) *RequestSLogger {
	return &RequestSLogger{id: id}
}

func (rl *RequestSLogger) WithService(service string) *RequestSLogger {
	rl.service = service
	return rl
}

func (rl *RequestSLogger) WithMethod(method string) *RequestSLogger {
	rl.method = method
	return rl
}

func (rl *RequestSLogger) WithResult(result any) *RequestSLogger {
	rl.result = result
	return rl
}

func (rl *RequestSLogger) WithStatus(status int) *RequestSLogger {
	rl.status = status
	return rl
}

func (rl *RequestSLogger) WithSince(since time.Duration) *RequestSLogger {
	rl.since = since
	return rl
}

func (rl *RequestSLogger) WithPath(path string) *RequestSLogger {
	rl.path = path
	return rl
}

func (rl *RequestSLogger) WithHost(host string) *RequestSLogger {
	rl.host = host
	return rl
}

func (rl RequestSLogger) Slog() []any {
	if rl.status >= http.StatusBadRequest {
		return []any{
			"x-request-id", rl.id.String(),
			"service", rl.service,
			"method", rl.method,
			"status", rl.status,
			"since", rl.since.String(),
			"host", rl.host,
			"path", rl.path,
			"error", rl.result,
		}
	}
	return []any{
		"x-request-id", rl.id.String(),
		"service", rl.service,
		"method", rl.method,
		"status", rl.status,
		"since", rl.since.String(),
		"host", rl.host,
		"path", rl.path,
	}
}

func (rl RequestSLogger) String() string {
	if slices.Contains([]int{http.StatusInternalServerError, http.StatusBadRequest}, rl.status) {
		return fmt.Sprintf(
			"%s | %s | %s | %s | %s | %s | %v",
			PadAndColorByStatus(rl.status, 7, rl.service),
			rl.id.String(),
			Pad(12, rl.since),
			PadAndColorByStatus(rl.status, 7, rl.method),
			Pad(12, rl.path),
			PadAndColorByStatus(rl.status, 0, rl.status),
			ColorByStatus(rl.status, rl.result),
		)
	}
	return fmt.Sprintf(
		"%s | %s | %s | %s | %s | %s",
		PadAndColorByStatus(rl.status, 7, rl.service),
		rl.id.String(),
		Pad(12, rl.since),
		PadAndColorByStatus(rl.status, 7, rl.method),
		Pad(12, rl.path),
		PadAndColorByStatus(rl.status, 0, rl.status),
	)
}
