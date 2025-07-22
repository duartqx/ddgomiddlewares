package logger

import (
	"cmp"
	"fmt"
	"math"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/duartqx/ddgomiddlewares/logger/colors"
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
			rl.padAndColor(7, rl.Method),
			rl.padAndColor(0, rl.Status),
			rl.pad(12, rl.Since),
			rl.Path,
			rl.getColored(rl.Result),
		)
	}
	return fmt.Sprintf(
		"| %s | %s |             | %s",
		rl.padAndColor(7, rl.Method),
		rl.padAndColor(0, rl.Status),
		rl.Path,
	)
}

func (rl RequestLogger) PanicString(err any) string {
	return fmt.Sprintf(
		"| %s | %s | %s | %s | %v",
		rl.padAndColor(7, rl.Method),
		rl.padAndColor(0, rl.Status),
		rl.pad(12, rl.Since),
		rl.Path,
		rl.getColored(cmp.Or(err, rl.Result)),
	)
}

func (rl RequestLogger) pad(padding int, value any) string {
	var (
		v string = fmt.Sprint(value)
		r int    = int(math.Max(float64(padding-len(v)), 0))
	)
	return v + strings.Repeat(" ", r)
}

func (rl RequestLogger) padAndColor(padding int, value any) string {
	if padding > 0 {
		return rl.getColored(rl.pad(padding, value))
	}
	return rl.getColored(value)
}

func (rl RequestLogger) getColored(value any) string {
	return colors.GetStatusColor(rl.Status, value)
}
