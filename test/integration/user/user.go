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
	return &nero.Schema{
		Pkg:        "repository",
		Collection: "users",
		Columns: []*nero.Column{
			nero.NewColumn("id", u.ID).
				StructField("ID").Ident().Auto(),
			nero.NewColumn("uid", u.UID).
				StructField("UID"),
			nero.NewColumn("email", u.Email),
			nero.NewColumn("name", u.Name),
			nero.NewColumn("age", u.Age),
			nero.NewColumn("group", u.Group).
				StructField("Group"),
			nero.NewColumn("kv", u.Kv),
			nero.NewColumn("tags", u.Tags),
			nero.NewColumn("updated_at", u.UpdatedAt).Nullable(),
			nero.NewColumn("created_at", u.CreatedAt).
				Auto(),
		},
	}
}
