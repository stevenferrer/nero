package internal

import (
	"reflect"

	"github.com/dave/jennifer/jen"
	"github.com/sf9v/mira"
)

// GetTypeC returns the jen.Code from a typ
func GetTypeC(typ *mira.Type) jen.Code {
	c := &jen.Statement{}
	nillable := typ.Nillable()

	// built-in types
	if typ.PkgPath() == "" {
		switch typ.T().Kind() {
		case reflect.Int:
			return c.Add(starc(nillable)).Int()
		case reflect.Int32:
			return c.Add(starc(nillable)).Int32()
		case reflect.Int64:
			return c.Add(starc(nillable)).Int64()
		case reflect.Uint:
			return c.Add(starc(nillable)).Uint()
		case reflect.Uint32:
			return c.Add(starc(nillable)).Uint32()
		case reflect.Uint64:
			return c.Add(starc(nillable)).Uint64()
		case reflect.Float32:
			return c.Add(starc(nillable)).Float32()
		case reflect.Float64:
			return c.Add(starc(nillable)).Float64()
		case reflect.String:
			return c.Add(starc(nillable)).String()
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

	return c.Add(starc(nillable)).Qual(typ.PkgPath(), typ.Name())
}

func starc(nillable bool) *jen.Statement {
	if nillable {
		return jen.Op("*")
	}

	return nil
}

// GetZeroValC returns the jen.Code zero value from a typ
func GetZeroValC(typ *mira.Type) jen.Code {
	if typ.Nillable() {
		return jen.Nil()
	}

	switch typ.T().Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return jen.Lit(0)
	case reflect.String:
		return jen.Lit("")
	}

	return jen.Qual(typ.PkgPath(), typ.Name()).Op("{").Op("}")
}
