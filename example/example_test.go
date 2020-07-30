package example_test

import (
	"log"
	"testing"

	"github.com/sf9v/nero/example"
	"github.com/sf9v/nero/gen"
)

func TestExample(t *testing.T) {
	err := gen.Generate(&example.User{}, "gen/user")
	if err != nil {
		log.Fatal(err)
	}
}
