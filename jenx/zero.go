package jenx

import (
	"reflect"

	"github.com/dave/jennifer/jen"
	"github.com/sf9v/mira"
)

func Zero(v interface{}) *jen.Statement {
	mt := mira.NewType(v)
	if mt.Kind() == mira.Numeric {
		return jen.Lit(0)
	}

	// built-in types
	if len(mt.PkgPath()) == 0 {
		switch mt.T().Kind() {
		case reflect.Bool:
			return jen.False()
		case reflect.String:
			return jen.Lit("")
		case reflect.Array:
			ev := reflect.New(mt.T().Elem()).Elem().Interface()
			return jen.Index(jen.Lit(mt.T().Len())).Add(Type(ev)).Values()
		}
		return jen.Nil()
	}

	// other types
	switch mt.T().Kind() {
	case reflect.String:
		return jen.Lit("")
	case reflect.Bool:
		return jen.False()
	case reflect.Struct:
		return jen.Qual(mt.PkgPath(), mt.Name()).Values()
	case reflect.Array:
		rt := resolveType(mt.T())
		if len(rt.Name()) == 0 {
			ev := reflect.New(mt.T().Elem()).Elem().Interface()
			return jen.Index(jen.Lit(mt.T().Len())).Add(Zero(ev))
		}

		return jen.Qual(mt.PkgPath(), mt.Name()).Values()
	}

	return jen.Nil()
}
