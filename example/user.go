package example

import (
	"time"

	"github.com/sf9v/nero"
)

// User is a basic example type
type User struct {
	ID        int64
	Name      string
	Group     string
	Age       int
	UpdatedAt *time.Time
	CreatedAt *time.Time
}

// Schema implements nero.Schemaer
func (u *User) Schema() *nero.Schema {
	return &nero.Schema{
		Pkg:        "user",
		Collection: "users",
		Columns: []*nero.Column{
			nero.NewColumn("id", u.ID).
				StructField("ID").Ident().Auto(),
			nero.NewColumn("name", u.Name),
			nero.NewColumn("group_res", u.Group).
				StructField("Group"),
			nero.NewColumn("updated_at", u.UpdatedAt),
			nero.NewColumn("created_at", u.CreatedAt).Auto(),
		},
	}
}
