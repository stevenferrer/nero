package jenx

import (
	"reflect"

	"github.com/dave/jennifer/jen"
	"github.com/sf9v/mira"
)

// Type value type statement
func Type(v interface{}) *jen.Statement {
	var (
		mt    = mira.NewType(v)
		rt    = resolveType(mt.T())
		isPtr = mt.T().Kind() == reflect.Ptr
		c     = new(jen.Statement)
	)
	// built-in types
	if len(mt.PkgPath()) == 0 {
		switch rt.Kind() {
		case reflect.Int:
			return c.Add(hasStar(isPtr)).Int()
		case reflect.Int8:
			return c.Add(hasStar(isPtr)).Int8()
		case reflect.Int16:
			return c.Add(hasStar(isPtr)).Int16()
		case reflect.Int32:
			return c.Add(hasStar(isPtr)).Int32()
		case reflect.Int64:
			return c.Add(hasStar(isPtr)).Int64()
		case reflect.Uint:
			return c.Add(hasStar(isPtr)).Uint()
		case reflect.Uint8:
			return c.Add(hasStar(isPtr)).Uint8()
		case reflect.Uint16:
			return c.Add(hasStar(isPtr)).Uint16()
		case reflect.Uint32:
			return c.Add(hasStar(isPtr)).Uint32()
		case reflect.Uint64:
			return c.Add(hasStar(isPtr)).Uint64()
		case reflect.Float32:
			return c.Add(hasStar(isPtr)).Float32()
		case reflect.Float64:
			return c.Add(hasStar(isPtr)).Float64()
		case reflect.Bool:
			return c.Add(hasStar(isPtr)).Bool()
		case reflect.String:
			return c.Add(hasStar(isPtr)).String()
		case reflect.Map:
			kv := reflect.New(mt.T().Key()).Elem().Interface()
			ev := reflect.New(mt.T().Elem()).Elem().Interface()
			return c.Map(Type(kv)).Add(Type(ev))
		case reflect.Array:
			ev := reflect.New(mt.T().Elem()).Elem().Interface()
			return c.Index(jen.Lit(mt.T().Len())).Add(Type(ev))
		case reflect.Slice:
			ev := reflect.New(mt.T().Elem()).Elem().Interface()
			return c.Index().Add(Type(ev))
		}
	}

	// external types
	switch rt.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64, reflect.Uint,
		reflect.Uint32, reflect.Uint64, reflect.Float32,
		reflect.Float64, reflect.String, reflect.Struct:
		return jen.Add(hasStar(isPtr)).Qual(mt.PkgPath(), mt.Name())
	case reflect.Map:
		if len(mt.Name()) == 0 {
			kt := reflect.New(mt.T().Key()).Elem().Interface()
			et := reflect.New(mt.T().Elem()).Elem().Interface()
			return c.Add(hasStar(isPtr)).Map(Type(kt)).Add(Type(et))
		}

		return jen.Add(hasStar(isPtr)).Qual(mt.PkgPath(), mt.Name())
	case reflect.Array:
		if len(rt.Name()) == 0 {
			ev := reflect.New(mt.T().Elem()).Elem().Interface()
			return c.Index(jen.Lit(mt.T().Len())).Add(Type(ev))
		}
		return c.Add(hasStar(isPtr)).Qual(mt.PkgPath(), mt.Name())
	case reflect.Slice:
		ev := reflect.New(mt.T().Elem()).Elem().Interface()
		return c.Index().Add(Type(ev))
	}

	return c
}

func hasStar(isPtr bool) *jen.Statement {
	if isPtr {
		return jen.Op("*")
	}
	return nil
}

func resolveType(t reflect.Type) reflect.Type {
	switch t.Kind() {
	case reflect.Ptr:
		return resolveType(t.Elem())
	}
	return t
}
