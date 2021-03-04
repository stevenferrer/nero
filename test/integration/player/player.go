package player

import (
	"time"

	"github.com/sf9v/nero"
)

// Player is a user
type Player struct {
	ID        string
	Email     string
	Name      string
	Age       int
	Race      Race
	Interests []string
	UpdatedAt *time.Time
	CreatedAt *time.Time
}

// Race is the player race
type Race string

// Factions
const (
	RaceHuman   Race = "human"
	RaceCharr   Race = "charr"
	RaceNorn    Race = "norn"
	RaceSylvari Race = "sylvari"
	RaceTitan   Race = "titan"
)

// Schema implements nero.Schemaer
func (u Player) Schema() *nero.Schema {
	return nero.NewSchemaBuilder(&u).
		PkgName("playerrepo").
		Collection("players").
		Identity(
			nero.NewColumnBuilder("id", u.ID).
				StructField("ID").Auto().Build(),
		).
		Columns(
			nero.NewColumnBuilder("email", u.Email).Build(),
			nero.NewColumnBuilder("name", u.Name).Build(),
			nero.NewColumnBuilder("age", u.Age).Build(),
			nero.NewColumnBuilder("race", u.Race).Build(),
			nero.NewColumnBuilder("interests", u.Interests).Build(),
			nero.NewColumnBuilder("updated_at", u.UpdatedAt).
				Optional().Build(),
			nero.NewColumnBuilder("created_at", u.CreatedAt).
				Auto().Build(),
		).
		Build()
}
