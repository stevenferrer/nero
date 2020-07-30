package internal

import (
	"reflect"
)

type Typ struct {
	Name    string
	V       interface{}
	T       reflect.Type
	PkgPath string
	Nillabe bool
}

func NewTyp(v interface{}) *Typ {
	t := reflect.TypeOf(v)

	rt := t
	for rt.Kind() == reflect.Ptr {
		rt = t.Elem()
	}

	return &Typ{
		T:       rt,
		V:       v,
		Name:    rt.Name(),
		PkgPath: rt.PkgPath(),
		Nillabe: nillable(t),
	}
}

func nillable(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Slice, reflect.Array, reflect.Ptr, reflect.Map:
		return true
	}
	return false
}
