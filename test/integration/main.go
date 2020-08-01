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
	outFiles, err := gen.Generate(new(user.User))
	checkErr(err)

	// create base directory
	basePath := path.Join("gen/user")
	err = os.MkdirAll(basePath, os.ModePerm)
	checkErr(err)

	// write files
	for _, outFile := range outFiles {
		filePath := path.Join(basePath, outFile.Name)
		file, err := os.Create(filePath)
		defer func(file *os.File) {
			checkErr(file.Close())
		}(file)
		checkErr(err)

		_, err = file.Write(outFile.Bytes())
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
