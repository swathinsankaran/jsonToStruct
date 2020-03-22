package main

import (
	"container/list"
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

	l := list.New()
	var value interface{}
	contents := readContents(fileName)
	err := json.Unmarshal(contents, &value)
	if err != nil {
		panic("Invalid JSON")
	}
	fmt.Println("FileName: ", fileName)
	fmt.Println("Generated struct: ")

	switch val := value.(type) {
	case map[string]interface{}:
		l.PushFront("type auto struct { ")
		parseJSONObjects(l, val, 2)
	case []interface{}:
		first = true
		l.PushFront("type auto []struct { ")
		_, _ = parseJSONArrays(l, val, 1, 2)
	default:
		panic("Invalid JSON provided.")
	}

	l.PushBack("}")
	print(l)
	fmt.Println("====================")
}

func print(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
