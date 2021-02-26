package example

import (
	"time"

	"github.com/sf9v/nero"
)

// Group is an example type that uses string ID
type Group struct {
	ID        string
	Name      string
	UpdatedAt *time.Time
	CreatedAt *time.Time
}

// Schema implements nero.Schamaer
func (g *Group) Schema() *nero.Schema {
	return nero.NewSchemaBuilder(g).
		PkgName("user").Collection("users").
		Identity(
			nero.NewColumnBuilder("id", g.ID).StructField("ID").
				Auto().Build(),
		).
		Columns(
			nero.NewColumnBuilder("name", g.Name).Build(),
			nero.NewColumnBuilder("updated_at", g.UpdatedAt).Build(),
			nero.NewColumnBuilder("created_at", g.CreatedAt).Auto().Build(),
		).Build()
}
