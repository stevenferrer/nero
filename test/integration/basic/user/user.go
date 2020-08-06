package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/sf9v/nero"
)

// User is a user
type User struct {
	ID        string
	UID       uuid.UUID
	Email     *string
	Name      *string
	Age       int
	Group     Group
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
		Pkg:        "user",
		Collection: "users",
		Columns: []*nero.Column{
			nero.NewColumn("id", u.ID).
				StructField("ID").Ident().Auto(),
			nero.NewColumn("uid", u.UID).
				StructField("UID"),
			nero.NewColumn("email", u.Email),
			nero.NewColumn("name", u.Name),
			nero.NewColumn("age", u.Age),
			nero.NewColumn("group_res", u.Group).
				StructField("Group"),
			nero.NewColumn("updated_at", u.UpdatedAt),
			nero.NewColumn("created_at", u.CreatedAt).
				Auto(),
		},
	}
}
