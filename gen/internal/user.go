package internal

import (
	"time"

	"github.com/stevenferrer/nero"
)

// User is a user model
type User struct {
	ID         int64
	Name       string
	Department string
	UpdatedAt  time.Time
	CreatedAt  time.Time
}

// Schema returns the schema for user model
func (u User) Schema() *nero.Schema {
	return nero.NewSchemaBuilder(&u).
		PkgName("userrepo").Table("users").
		Identity(
			nero.NewFieldBuilder("id", u.ID).
				StructField("ID").Auto().Build(),
		).
		Fields(
			nero.NewFieldBuilder("name", u.Name).
				Build(),
			nero.NewFieldBuilder("department", u.Department).
				Build(),
			nero.NewFieldBuilder("updated_at", u.UpdatedAt).
				Optional().Build(),
			nero.NewFieldBuilder("created_at", u.CreatedAt).
				Auto().Build(),
		).
		Templates(nero.NewPostgresTemplate()).
		Build()
}
