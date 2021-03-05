package internal

import (
	"time"

	"github.com/sf9v/nero"
)

// User is a user model
type User struct {
	ID         int64
	Name       string
	Department string
	UpdatedAt  *time.Time
	CreatedAt  *time.Time
}

// Schema returns the schema for user model
func (u User) Schema() *nero.Schema {
	return nero.NewSchemaBuilder(&u).
		PkgName("userrepo").Collection("users").
		Identity(
			nero.NewColumnBuilder("id", u.ID).StructField("ID").
				Auto().Build(),
		).
		Columns(
			nero.NewColumnBuilder("name", u.Name).Build(),
			nero.NewColumnBuilder("department", u.Department).Build(),
			nero.NewColumnBuilder("updated_at", u.UpdatedAt).
				Optional().Comparable().Build(),
			nero.NewColumnBuilder("created_at", u.CreatedAt).
				Auto().Build(),
		).
		Templates(
			nero.NewPostgresTemplate().WithFilename("postgres.go"),
		).
		Build()
}
