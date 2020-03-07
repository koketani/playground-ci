package greet

import (
	"fmt"
	"strings"

	"rsc.io/quote"
)

// Greet ...
func Greet(toWhom string) string {
	return fmt.Sprintf("%s, %s", toWhom, strings.ToLower(quote.Hello()))
}
