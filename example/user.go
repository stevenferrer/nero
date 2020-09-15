package example

import (
	"time"

	"github.com/google/uuid"
	"github.com/sf9v/nero"
)

// User is a basic example type
type User struct {
	ID           int64
	UUID         uuid.UUID
	Name         string
	Group        string
	Age          int
	IsRegistered bool
	Tags         [10]string
	Empty        struct{}
	UpdatedAt    *time.Time
	CreatedAt    *time.Time
}

// Schema implements nero.Schemaer
func (u *User) Schema() *nero.Schema {
	return &nero.Schema{
		Pkg:        "user",
		Collection: "users",
		Columns: []*nero.Column{
			nero.NewColumn("id", u.ID).
				StructField("ID").Ident().Auto(),
			nero.NewColumn("uuid", u.UUID).StructField("UUID"),
			nero.NewColumn("name", u.Name),
			nero.NewColumn("group_res", u.Group).
				StructField("Group"),
			nero.NewColumn("age", u.Age),
			nero.NewColumn("is_registered", u.IsRegistered),
			nero.NewColumn("tags", u.Tags),
			nero.NewColumn("empty", u.Empty),
			nero.NewColumn("updated_at", u.UpdatedAt),
			nero.NewColumn("created_at", u.CreatedAt).Auto(),
		},
	}
}
