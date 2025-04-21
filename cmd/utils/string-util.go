package utils

import "strings"

func CapitalizeFirstLetter(word string) string {
	return strings.ToUpper(word[:1]) + word[1:]
}
