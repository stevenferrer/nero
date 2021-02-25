package customtypes

import (
	"github.com/segmentio/ksuid"
	"github.com/sf9v/nero"
)

// Custom demonstrates the use of many different field types
type Custom struct {
	ID             int64
	S              string
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
	KSID           ksuid.KSUID
	NullColumn     *string
}

// Item is an example struct embedded in Custom struct
type Item struct {
	Name string
}

// Schema implements nero.Schemaer
func (c *Custom) Schema() *nero.Schema {
	return nero.NewSchemaBuilder().
		PkgName("user").Collection("users").
		Columns(
			nero.NewColumnBuilder("id", c.ID).
				StructField("ID").Identity().Auto().Build(),
			nero.NewColumnBuilder("ksid", c.KSID).StructField("KSID").Build(),
			nero.NewColumnBuilder("s", c.S).Build(),
			nero.NewColumnBuilder("map_str_str", c.MapStrStr).Build(),
			nero.NewColumnBuilder("map_str_ptr_str", c.MapStrPtrStr).Build(),
			nero.NewColumnBuilder("map_int64_str", c.MapInt64Str).Build(),
			nero.NewColumnBuilder("map_int64_ptr_str", c.MapInt64PtrStr).Build(),
			nero.NewColumnBuilder("map_str_item", c.MapStrItem).Build(),
			nero.NewColumnBuilder("map_str_ptr_item", c.MapStrPtrItem).Build(),
			nero.NewColumnBuilder("item", c.Item).Build(),
			nero.NewColumnBuilder("ptr_item", c.PtrItem).Build(),
			nero.NewColumnBuilder("items", c.Items).Build(),
			nero.NewColumnBuilder("ptr_items", c.PtrItems).Build(),
			nero.NewColumnBuilder("null_column", c.NullColumn).Build(),
		).
		Build()
}
