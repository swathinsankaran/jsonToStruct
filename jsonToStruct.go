package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"log"

	"github.com/iancoleman/strcase"
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
		_, _ = jts.parseJSONArrays(value.([]interface{}), 1, 1, isArray)
		return
	}
	jts.parseJSONObjects(value.(map[string]interface{}), 1, isArray)
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

// parseJSONObjects method parses the JSON objects and constructs the final struct.
func (jts *jsonToStruct) parseJSONObjects(e map[string]interface{}, tabCount int, isFirst bool) {

	for key, value := range e {
		switch val := value.(type) {
		case string:
			jts.generatedStruct.PushBack(fmt.Sprintf("%s%-10s%-10s%-10s", printTabs(tabCount), strcase.ToCamel(key), "string", "`json:\""+strcase.ToSnake(key)+"\"`"))
		case []interface{}:
			var tabs int
			if isFirst && tabCount == 1 {
				jts.generatedStruct.PushBack(fmt.Sprintf("%-10s[]struct %-10s", strcase.ToCamel(key), "{ "))
				isFirst = false
				tabs = tabCount
				tabCount += 1
			}
			t, length := jts.parseJSONArrays(val, 1, tabCount, isFirst)
			if len(t) != 0 {
				dataType := getType(t, length)
				jts.generatedStruct.PushBack(fmt.Sprintf("%s%-10s%-10s%-10s", printTabs(tabCount), strcase.ToCamel(key), dataType, "`json:\""+strcase.ToSnake(key)+"\"`"))
			}
			if tabs == 2 {
				jts.generatedStruct.PushBack(fmt.Sprintf("%s%-10s%-10s", printTabs(tabs), "}", "`json:\""+strcase.ToSnake(key)+"\"`"))
			}
		case float64:
			jts.generatedStruct.PushBack(fmt.Sprintf("%s%-10s%-10s%-10s", printTabs(tabCount), strcase.ToCamel(key), "int", "`json:\""+strcase.ToSnake(key)+"\"`"))
		case map[string]interface{}:
			jts.generatedStruct.PushBack(fmt.Sprintf("%s%-10sstruct %-10s", printTabs(tabCount), strcase.ToCamel(key), "{ "))
			jts.parseJSONObjects(val, tabCount+1, isFirst)
			jts.generatedStruct.PushBack(fmt.Sprintf("%s%s %s", printTabs(tabCount), "}", "`json:\""+strcase.ToSnake(key)+"\"`"))
		}
	}
}

// parseJSONArrays method parses the JSON arrys and calls the JSON object parser function for further parsing.
func (jts *jsonToStruct) parseJSONArrays(e []interface{}, level, tabCount int, isFirst bool) ([]string, int) {

	var dataTypes []string
	var counter map[string]int
	var lastKey string
	counter = make(map[string]int)
	for i, value := range e {
		switch val := value.(type) {
		case float64:
			dataTypes = append(dataTypes, "int")
		case string:
			dataTypes = append(dataTypes, "string")
		case []interface{}:
			if i == 0 {
				level++
			}
			dataTypes, level = jts.parseJSONArrays(val, level, tabCount, isFirst)
		case map[string]interface{}:
			for key := range val {
				counter[key]++
				lastKey = key
			}
			if counter[lastKey] > 1 {
				continue
			}
			jts.parseJSONObjects(val, tabCount, isFirst)
		}
	}
	return dataTypes, level
}
