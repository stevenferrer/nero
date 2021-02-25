package example

import "github.com/sf9v/nero"

// NoPreds is an example struct which has fields that don't have predicates
type NoPreds struct {
	ID       int64
	M        map[string]string
	A        [1]string
	S        []string
	MyStruct *MyStruct
}

// MyStruct is an example struct
type MyStruct struct {
	S string
}

// Schema implements nero.Schemaer
func (n *NoPreds) Schema() *nero.Schema {
	return nero.NewSchemaBuilder().
		PkgName("user").Collection("users").
		Columns(
			nero.NewColumnBuilder("id", n.ID).
				StructField("ID").Identity().Auto().Build(),
			nero.NewColumnBuilder("m", n.M).Build(),
			nero.NewColumnBuilder("a", n.A).Build(),
			nero.NewColumnBuilder("s", n.S).Build(),
		).Build()
}
