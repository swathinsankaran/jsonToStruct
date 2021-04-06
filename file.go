package main

import (
	"io/ioutil"
	"os"
)

type file struct {
	filePath string
	contents []byte
}

func newFile(filePath string) *file {
	f := &file{}
	f.filePath = filePath
	f.Read()
	return f
}

func (f *file) Read() {
	fileReader, err := os.Open(f.filePath)
	checkErr(err)
	defer fileReader.Close()

	f.contents, err = ioutil.ReadAll(fileReader)
	checkErr(err)
}
