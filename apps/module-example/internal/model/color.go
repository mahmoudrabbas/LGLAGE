package model

import (
	"fmt"
	"strconv"
	"strings"
)

type Color struct {
	value int
}

var (
	Black     = Color{value: 30}
	Red       = Color{value: 31}
	Green     = Color{value: 32}
	Bold      = Color{value: 1}
	Underline = Color{value: 4}
)

func Text(text string, props ...Color) string {

	if len(props) == 0 {
		return text
	}

	var codes []string

	for _, color := range props {
		codes = append(codes, strconv.Itoa(color.value))
	}

	return fmt.Sprintf("\033[%sm%s\033[0m", strings.Join(codes, ";"), text)
}
