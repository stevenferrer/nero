package player

import (
	"time"

	"github.com/sf9v/nero"
)

// Player is a plaer
type Player struct {
	ID        string
	Email     string
	Name      string
	Age       int
	Race      Race
	UpdatedAt *time.Time
	CreatedAt *time.Time
}

// Race is the player race
type Race string

// List of player race
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
		Table("players").
		Identity(nero.NewFieldBuilder("id", p.ID).
			StructField("ID").Auto().Build()).
		Fields(
			nero.NewFieldBuilder("email", p.Email).Build(),
			nero.NewFieldBuilder("name", p.Name).Build(),
			nero.NewFieldBuilder("age", p.Age).Build(),
			nero.NewFieldBuilder("race", p.Race).Build(),
			nero.NewFieldBuilder("updated_at", p.UpdatedAt).
				Optional().Build(),
			nero.NewFieldBuilder("created_at", p.CreatedAt).
				Auto().Build(),
		).
		Templates(
			nero.NewPostgresTemplate(),
			nero.NewSQLiteTemplate(),
		).
		Build()
}
