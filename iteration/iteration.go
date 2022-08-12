package iteration

import "strings"
func Repeat(character string, iterations  int) (repeat string) {

	repeat = strings.Repeat(character, iterations)
	return repeat
}

