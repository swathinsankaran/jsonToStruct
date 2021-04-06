package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"log"
)

type jsonToStruct struct {
	file            *file
	generatedStruct *list.List
}

func (jts *jsonToStruct) Convert() {

	var value interface{}
	err := json.Unmarshal(jts.file.contents, &value)
	checkErr(err)

	var isArray bool

	switch val := value.(type) {
	case map[string]interface{}:
	case []interface{}:
		if len(val) == 0 {
			log.Fatal("Empty JSON provided.")
		}
		isArray = true
	default:
		log.Fatal("Invalid or JSON provided.")
	}

	jts.header(isArray)
	defer jts.footer()

	if isArray {
		_, _ = parseJSONArrays(jts.generatedStruct, value.([]interface{}), 1, 1, isArray)
		return
	}
	parseJSONObjects(jts.generatedStruct, value.(map[string]interface{}), 1, isArray)
}

func (jts *jsonToStruct) Print() {
	log.Println("FileName: ", jts.file.filePath)
	log.Println("Generated struct: ")
	listNode := jts.generatedStruct.Front()
	for listNode != nil {
		fmt.Println(listNode.Value)
		listNode = listNode.Next()
	}
}

func (jts *jsonToStruct) header(isArray bool) {
	const header string = "type T %sstruct { "
	if isArray {
		jts.generatedStruct.PushFront(fmt.Sprintf(header, "[]"))
		return
	}
	jts.generatedStruct.PushFront(fmt.Sprintf(header, ""))
}

func (jts *jsonToStruct) footer() {
	jts.generatedStruct.PushBack("}")
}
