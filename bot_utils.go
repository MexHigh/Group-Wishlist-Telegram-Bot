package main

import (
	"fmt"
	"unicode"
)

// capError capitalizes the error
// message of an error type
func beautifulError(err error) string {
	str := []rune(err.Error())
	str[0] = unicode.ToUpper(str[0])
	return fmt.Sprintf("Error: %s", string(str))
}
