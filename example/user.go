package example

import (
	"time"

	"github.com/google/uuid"

	"github.com/sf9v/nero"
	"github.com/sf9v/nero/template"
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
	return nero.NewSchemaBuilder().
		PkgName("user").Collection("users").
		Columns(
			nero.NewColumnBuilder("id", u.ID).
				StructField("ID").Identity().Auto().Build(),
			nero.NewColumnBuilder("uuid", u.UUID).StructField("UUID").Build(),
			nero.NewColumnBuilder("name", u.Name).Build(),
			nero.NewColumnBuilder("group_res", u.Group).
				StructField("Group").Build(),
			nero.NewColumnBuilder("age", u.Age).Build(),
			nero.NewColumnBuilder("is_registered", u.IsRegistered).Build(),
			nero.NewColumnBuilder("tags", u.Tags).Build(),
			nero.NewColumnBuilder("empty", u.Empty).Build(),
			nero.NewColumnBuilder("updated_at", u.UpdatedAt).ColumnComparable().Build(),
			nero.NewColumnBuilder("created_at", u.CreatedAt).Auto().Build(),
		).
		Templates(template.NewPostgresTemplate().WithFilename("postgres.go")).
		Build()
}
