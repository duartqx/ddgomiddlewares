package logger

import (
	"fmt"
	"math"
	"strings"

	"github.com/duartqx/ddgomiddlewares/logger/colors"
)

func Pad(padding int, value any) string {
	var (
		v string = fmt.Sprint(value)
		r int    = int(math.Max(float64(padding-len(v)), 0))
	)
	return v + strings.Repeat(" ", r)
}

func PadAndColorByStatus(status, padding int, value any) string {
	if padding > 0 {
		return ColorByStatus(status, Pad(padding, value))
	}
	return ColorByStatus(status, value)
}

func ColorByStatus(status int, value any) string {
	return colors.GetStatusColor(status, value)
}
