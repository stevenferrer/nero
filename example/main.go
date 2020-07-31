package main

import (
	"log"
	"os"
	"path"

	"github.com/sf9v/nero/example/user"
	"github.com/sf9v/nero/gen"
)

func main() {
	outFiles, err := gen.Generate(new(user.User))
	checkErr(err)

	// create base directory
	basePath := path.Join("gen/user")
	err = os.MkdirAll(basePath, os.ModePerm)
	checkErr(err)

	// write files
	for _, outFile := range outFiles {
		filePath := path.Join(basePath, outFile.Name)
		f, err := os.Create(filePath)
		checkErr(err)

		_, err = f.Write(outFile.Bytes())
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
