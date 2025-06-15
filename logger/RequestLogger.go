package logger

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/duartqx/ddgomiddlewares/logger/colors"
)

type RequestLogger struct {
	Method string
	Status int
	Since  time.Duration
	Path   string
}

func NewRequestLogger() *RequestLogger {
	return &RequestLogger{}
}

func (rl RequestLogger) String() string {
	return fmt.Sprintf(
		"| %s | %s | %s | %s",
		rl.padAndColor(7, rl.Method),
		rl.padAndColor(0, rl.Status),
		rl.pad(12, rl.Since),
		rl.Path,
	)
}

func (rl RequestLogger) PanicString(err any) string {
	return fmt.Sprintf(
		"| %s | %s |              | %s | %s",
		rl.padAndColor(7, rl.Method),
		rl.padAndColor(0, rl.Status),
		rl.Path,
		rl.getColored(err),
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
	var colored string
	switch {
	case rl.Status >= 100 && rl.Status < 200:
		colored, _ = colors.ColorIt(colors.Cyan, value)
	case rl.Status >= 200 && rl.Status < 300:
		colored, _ = colors.ColorIt(colors.Green, value)
	case rl.Status >= 300 && rl.Status < 400:
		colored, _ = colors.ColorIt(colors.Magenta, value)
	case rl.Status >= 400 && rl.Status < 500:
		colored, _ = colors.ColorIt(colors.Yellow, value)
	default:
		colored, _ = colors.ColorIt(colors.Red, value)
	}
	return colored
}
