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
func (np *NoPreds) Schema() *nero.Schema {
	return nero.NewSchemaBuilder(np).
		PkgName("user").Collection("users").
		Identity(
			nero.NewColumnBuilder("id", np.ID).
				StructField("ID").Auto().Build(),
		).
		Columns(
			nero.NewColumnBuilder("m", np.M).Build(),
			nero.NewColumnBuilder("a", np.A).Build(),
			nero.NewColumnBuilder("s", np.S).Build(),
		).Build()
}
