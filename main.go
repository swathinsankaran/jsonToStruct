package main

import (
	"errors"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		checkErr(errors.New("Usage: ./jsonToStruct <file1> <file2>..."))
	}

	for _, filePath := range os.Args[1:] {
		process(filePath)
	}
}

func process(filePath string) {
	c := newConvertor("JSONToStruct", filePath)
	c.Convert()
	c.Print()
}
