package example

import (
	"time"

	"github.com/sf9v/nero"
)

type User struct {
	ID        int64
	Email     string
	Name      string
	UpdatedAt *time.Time
	CreatedAt *time.Time
}

func (u *User) Schema() *nero.Schema {
	return &nero.Schema{
		Pkg:        "user",
		Collection: "users",
		Columns: nero.Columns{
			nero.NewColumn("id", u.ID).
				Field("ID").Ident().Auto(),
			nero.NewColumn("email", u.Email),
			nero.NewColumn("name", u.Name),
			nero.NewColumn("updated_at", u.UpdatedAt),
			nero.NewColumn("created_at", u.CreatedAt).
				Auto(),
		},
	}
}
