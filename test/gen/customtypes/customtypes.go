package customtypes

import (
	"github.com/stevenferrer/nero"
)

// UUID is a uuid type
type UUID [16]byte

// Custom demonstrates the use of many different field types
type Custom struct {
	ID             int64
	UUID           UUID
	Str            string
	MapStrStr      map[string]string
	MapStrPtrStr   map[string]*string
	MapInt64Str    map[int64]string
	MapInt64PtrStr map[int64]*string
	MapStrItem     map[string]Item
	MapStrPtrItem  map[string]*Item
	Item           Item
	PtrItem        *Item
	Items          []Item
	PtrItems       []*Item
	NullColumn     *string
}

// Item is an example struct embedded in Custom struct
//
// Note: Custom types like these must implement ValueScanner
type Item struct {
	Name string
}

// Schema implements nero.Schemaer
func (c Custom) Schema() *nero.Schema {
	return nero.NewSchemaBuilder(&c).
		PkgName("user").Table("users").
		Identity(
			nero.NewFieldBuilder("id", c.ID).
				StructField("ID").Auto().Build(),
		).
		Fields(
			nero.NewFieldBuilder("uuid", c.UUID).StructField("UUID").Build(),
			nero.NewFieldBuilder("str", c.Str).Build(),
			nero.NewFieldBuilder("map_str_str", c.MapStrStr).Build(),
			nero.NewFieldBuilder("map_str_ptr_str", c.MapStrPtrStr).Build(),
			nero.NewFieldBuilder("map_int64_str", c.MapInt64Str).Build(),
			nero.NewFieldBuilder("map_int64_ptr_str", c.MapInt64PtrStr).Build(),
			nero.NewFieldBuilder("map_str_item", c.MapStrItem).Build(),
			nero.NewFieldBuilder("map_str_ptr_item", c.MapStrPtrItem).Build(),
			nero.NewFieldBuilder("item", c.Item).Build(),
			nero.NewFieldBuilder("ptr_item", c.PtrItem).Build(),
			nero.NewFieldBuilder("items", c.Items).Build(),
			nero.NewFieldBuilder("ptr_items", c.PtrItems).Build(),
			nero.NewFieldBuilder("null_column", c.NullColumn).Build(),
		).Build()
}
