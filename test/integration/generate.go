package main

import (
	"log"
	"os"
	"path"

	"github.com/sf9v/nero/gen"
	"github.com/sf9v/nero/test/integration/user"
)

func main() {
	// generate
	files, err := gen.Generate(new(user.User))
	checkErr(err)

	// create base directory
	basePath := path.Join("repository")
	err = os.MkdirAll(basePath, os.ModePerm)
	checkErr(err)

	for _, f := range files {
		err = f.Render(basePath)
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
