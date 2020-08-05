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
	return &nero.Schema{
		Pkg:        "user",
		Collection: "users",
		Columns: []*nero.Column{
			nero.NewColumn("id", g.ID).StructField("ID").Ident().Auto(),
			nero.NewColumn("name", g.Name),
			nero.NewColumn("updated_at", g.UpdatedAt),
			nero.NewColumn("created_at", g.CreatedAt).Auto(),
		},
	}
}
