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
