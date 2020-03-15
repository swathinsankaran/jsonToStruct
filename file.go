package main

import (
	"io/ioutil"
	"log"
	"os"
)

// readContents reads the file and returns the content of the file.
func readContents(fileName string) []byte {
	fileReader, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fileReader.Close()
	content, err := ioutil.ReadAll(fileReader)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
