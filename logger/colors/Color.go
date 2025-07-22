package colors

import (
	"fmt"
	"slices"
)

type Color string

const (
	Blue    Color = "\033[34m"
	Cyan    Color = "\033[36m"
	Green   Color = "\033[32m"
	Magenta Color = "\033[35m"
	Red     Color = "\033[31m"
	Yellow  Color = "\033[33m"

	Reset Color = "\033[0m"
)

var clrs []Color = []Color{Blue, Cyan, Green, Magenta, Red, Yellow}

func ColorIt(color Color, value any) (string, error) {
	if !slices.Contains(clrs, color) {
		return "", fmt.Errorf("Color not recognized!")
	}
	return fmt.Sprintf("%s%v%s", color, value, Reset), nil
}

func GetStatusColor(status int, value any) string {
	var colored string
	switch {
	case status >= 100 && status < 200:
		colored, _ = ColorIt(Cyan, value)
	case status >= 200 && status < 300:
		colored, _ = ColorIt(Green, value)
	case status >= 300 && status < 400:
		colored, _ = ColorIt(Magenta, value)
	case status >= 400 && status < 500:
		colored, _ = ColorIt(Yellow, value)
	default:
		colored, _ = ColorIt(Red, value)
	}
	return colored
}
