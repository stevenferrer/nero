package jenx

import (
	"github.com/dave/jennifer/jen"
)

// Dotln dot operator followed by a newline
func Dotln(name string) *jen.Statement {
	return jen.Op(".").Line().Id(name)
}
