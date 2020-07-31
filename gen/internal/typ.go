package internal

import (
	"reflect"
)

// Typ is a type info
type Typ struct {
	Name    string
	V       interface{}
	T       reflect.Type
	PkgPath string
	Nillabe bool
}

// NewTyp returns a type infor from v
func NewTyp(v interface{}) *Typ {
	t := reflect.TypeOf(v)
	rt := rtype(t)
	return &Typ{
		T:       rt,
		V:       v,
		Name:    rt.Name(),
		PkgPath: rt.PkgPath(),
		Nillabe: nillable(t),
	}
}

func rtype(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t
}

func nillable(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Slice, reflect.Array,
		reflect.Ptr, reflect.Map:
		return true
	}
	return false
}
