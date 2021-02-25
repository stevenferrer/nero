package user

import (
	"time"

	"github.com/segmentio/ksuid"
	"github.com/sf9v/nero"
	"github.com/sf9v/nero/example"
)

// User is a user
type User struct {
	ID        string
	UID       ksuid.KSUID
	Email     string
	Name      string
	Age       int
	Group     Group
	Kv        example.Map
	Tags      []string
	UpdatedAt *time.Time
	CreatedAt *time.Time
}

// Group is a group
type Group string

// Groups
const (
	Human   Group = "human"
	Charr   Group = "charr"
	Norn    Group = "norn"
	Sylvari Group = "sylvari"
	Outcast Group = "outcast"
)

// Schema implements nero.Schemaer
func (u *User) Schema() *nero.Schema {
	return nero.NewSchemaBuilder().
		PkgName("repository").Collection("users").
		Columns(
			nero.NewColumnBuilder("id", u.ID).
				StructField("ID").Identity().Auto().Build(),
			nero.NewColumnBuilder("uid", u.UID).
				StructField("UID").Build(),
			nero.NewColumnBuilder("email", u.Email).Build(),
			nero.NewColumnBuilder("name", u.Name).Build(),
			nero.NewColumnBuilder("age", u.Age).Build(),
			nero.NewColumnBuilder("group", u.Group).
				StructField("Group").Build(),
			nero.NewColumnBuilder("kv", u.Kv).Build(),
			nero.NewColumnBuilder("tags", u.Tags).Build(),
			nero.NewColumnBuilder("updated_at", u.UpdatedAt).Build(),
			nero.NewColumnBuilder("created_at", u.CreatedAt).
				Auto().Build(),
		).
		Build()
}
