package helper

import "strings"

func CamelStyle(o string) string {
	words := strings.Title(o)
	return strings.Join(strings.Fields(words), "")
}