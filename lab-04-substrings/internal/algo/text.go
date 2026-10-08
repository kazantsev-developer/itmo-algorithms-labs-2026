package algo

import (
	"strconv"
	"strings"
)

func BuildText(numbers []int) string {
	var sb strings.Builder
	for _, n := range numbers {
		sb.WriteString(strconv.Itoa(n))
	}
	return sb.String()
}
