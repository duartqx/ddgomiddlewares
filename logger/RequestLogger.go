package logger

import (
	"cmp"
	"fmt"
	"net/http"
	"slices"
	"time"
)

type RequestLogger struct {
	Method string
	Result any
	Status int
	Since  time.Duration
	Path   string
}

func NewRequestLogger() *RequestLogger {
	return &RequestLogger{}
}

func (rl RequestLogger) String() string {
	if slices.Contains([]int{http.StatusInternalServerError, http.StatusBadRequest}, rl.Status) {
		return fmt.Sprintf(
			"| %s | %s | %s | %s | %v",
			PadAndColorByStatus(rl.Status, 7, rl.Method),
			PadAndColorByStatus(rl.Status, 0, rl.Status),
			Pad(12, rl.Since),
			rl.Path,
			ColorByStatus(rl.Status, rl.Result),
		)
	}
	return fmt.Sprintf(
		"| %s | %s |             | %s",
		PadAndColorByStatus(rl.Status, 7, rl.Method),
		PadAndColorByStatus(rl.Status, 0, rl.Status),
		rl.Path,
	)
}

func (rl RequestLogger) PanicString(err any) string {
	return fmt.Sprintf(
		"| %s | %s | %s | %s | %v",
		PadAndColorByStatus(rl.Status, 7, rl.Method),
		PadAndColorByStatus(rl.Status, 0, rl.Status),
		Pad(12, rl.Since),
		rl.Path,
		ColorByStatus(rl.Status, cmp.Or(err, rl.Result)),
	)
}
