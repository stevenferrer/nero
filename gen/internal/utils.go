package internal

import (
	"reflect"

	"github.com/dave/jennifer/jen"
	"github.com/sf9v/mira"
)

// GetTypeC returns the jen.Code from a typ
func GetTypeC(typ *mira.Type) jen.Code {
	c := &jen.Statement{}
	nillable := typ.IsNillable()

	// built-in types
	if typ.PkgPath() == "" {
		switch typ.T().Kind() {
		case reflect.Int:
			return c.Add(star(nillable)).Int()
		case reflect.Int32:
			return c.Add(star(nillable)).Int32()
		case reflect.Int64:
			return c.Add(star(nillable)).Int64()
		case reflect.Uint:
			return c.Add(star(nillable)).Uint()
		case reflect.Uint32:
			return c.Add(star(nillable)).Uint32()
		case reflect.Uint64:
			return c.Add(star(nillable)).Uint64()
		case reflect.Float32:
			return c.Add(star(nillable)).Float32()
		case reflect.Float64:
			return c.Add(star(nillable)).Float64()
		case reflect.Bool:
			return c.Add(star(nillable)).Bool()
		case reflect.String:
			return c.Add(star(nillable)).String()
		case reflect.Map:
			// get key type
			kt := mira.NewType(reflect.New(typ.T().Key()).Elem().Interface())
			// get element type
			et := mira.NewType(reflect.New(typ.T().Elem()).Elem().Interface())
			return c.Map(GetTypeC(kt)).Add(GetTypeC(et))
		case reflect.Slice:
			et := mira.NewType(reflect.New(typ.T().Elem()).Elem().Interface())
			return c.Index().Add(GetTypeC(et))
		}
	}

	return c.Add(star(nillable)).Qual(typ.PkgPath(), typ.Name())
}

func star(nillable bool) *jen.Statement {
	if nillable {
		return jen.Op("*")
	}

	return nil
}

// GetZeroValC returns the jen.Code zero value from a typ
func GetZeroValC(typ *mira.Type) jen.Code {
	if typ.IsNillable() {
		return jen.Nil()
	}

	if typ.IsNumeric() {
		return jen.Lit(0)
	}

	switch typ.T().Kind() {
	case reflect.String:
		return jen.Lit("")
	case reflect.Bool:
		return jen.False()
	}

	return jen.Qual(typ.PkgPath(), typ.Name()).Op("{").Op("}")
}
