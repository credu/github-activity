package utils

import "strings"

func CapitalizeFirstLetter(word string) string {
	return strings.ToUpper(string(word[0])) + word[1:]
}
