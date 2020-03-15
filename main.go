package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var first bool

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./jsonToGo <file1> <file2>...")
		os.Exit(1)
	}
	for _, fileName := range os.Args[1:] {
		processFile(fileName)
	}
}

// processFile function reads the content of the file and calls the parser functions.
func processFile(fileName string) {
	var value interface{}
	contents := readContents(fileName)
	err := json.Unmarshal(contents, &value)
	if err != nil {
		panic("Invalid JSON")
	}
	fmt.Println("FileName: ", fileName)
	fmt.Println("Generated struct: ")
	var finalStruct string
	val, ok := value.(map[string]interface{})
	if ok {
		finalStruct += "type auto struct { "
		parseJSONObjects(val, 2, &finalStruct)
		finalStruct += "}"
		fmt.Println(finalStruct)
		fmt.Println("====================")
		return
	}

	v, ok := value.([]interface{})
	if ok {
		first = true
		finalStruct += "type auto []struct { "
		_, _ = parseJSONArrays(v, 1, 2, &finalStruct)
		finalStruct += "}"
		fmt.Println(finalStruct)
		fmt.Println("====================")
	}
}
