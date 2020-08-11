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
	return &nero.Schema{
		Pkg:        "user",
		Collection: "users",
		Columns: []*nero.Column{
			nero.NewColumn("id", n.ID).
				StructField("ID").Ident().Auto(),
			nero.NewColumn("m", n.M),
			nero.NewColumn("a", n.A),
			nero.NewColumn("s", n.S),
		},
	}
}
