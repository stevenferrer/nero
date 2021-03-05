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
func (p Player) Schema() *nero.Schema {
	return nero.NewSchemaBuilder(&p).
		PkgName("playerrepo").
		Collection("players").
		Identity(
			nero.NewColumnBuilder("id", p.ID).
				StructField("ID").Auto().Build(),
		).
		Columns(
			nero.NewColumnBuilder("email", p.Email).Build(),
			nero.NewColumnBuilder("name", p.Name).Build(),
			nero.NewColumnBuilder("age", p.Age).Build(),
			nero.NewColumnBuilder("race", p.Race).Build(),
			nero.NewColumnBuilder("interests", p.Interests).Build(),
			nero.NewColumnBuilder("updated_at", p.UpdatedAt).
				Optional().Build(),
			nero.NewColumnBuilder("created_at", p.CreatedAt).
				Auto().Build(),
		).
		Build()
}
