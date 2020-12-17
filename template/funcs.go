package template

import (
	"fmt"
	"reflect"

	"github.com/jinzhu/inflection"
	"github.com/sf9v/mira"
	stringsx "github.com/sf9v/nero/x/strings"
)

func typeFunc(v interface{}) string {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Ptr {
		return fmt.Sprintf("%T", v)
	}

	ev := reflect.New(resolveType(t)).Elem().Interface()
	return fmt.Sprintf("%T", ev)
}

func resolveType(t reflect.Type) reflect.Type {
	switch t.Kind() {
	case reflect.Ptr:
		return resolveType(t.Elem())
	}
	return t
}

func zeroFunc(v interface{}) string {
	mt := mira.NewType(v)

	if mt.IsNillable() {
		return "nil"
	}

	if mt.Kind() == mira.Numeric {
		return "0"
	}

	switch mt.T().Kind() {
	case reflect.Bool:
		return "false"
	case reflect.Struct:
		return fmt.Sprintf("(%T{})", v)
	case reflect.Array:
		if len(mt.Name()) == 0 {
			ev := reflect.New(mt.T().Elem()).Elem().Interface()
			return fmt.Sprintf("[%d]%T{}", mt.T().Len(), ev)
		}

		return fmt.Sprintf("(%T{})", v)
	}

	return "\"\""

}

func pluralFunc(s string) string {
	return inflection.Plural(s)
}

func lowerCamelFunc(s string) string {
	return stringsx.ToLowerCamel(s)
}
