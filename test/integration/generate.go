package main

import (
	"log"
	"os"
	"path"

	"github.com/stevenferrer/nero/gen"
	"github.com/stevenferrer/nero/test/integration/playerpkg"
)

func main() {
	// generate
	p := playerpkg.Player{}
	files, err := gen.Generate(p.Schema())
	checkErr(err)

	// create base directory
	basePath := path.Join("playerrepo")
	err = os.MkdirAll(basePath, os.ModePerm)
	checkErr(err)

	for _, file := range files {
		err = file.Render(basePath)
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
