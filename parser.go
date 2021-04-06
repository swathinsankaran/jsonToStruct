package main

import (
	"container/list"
	"fmt"

	"github.com/iancoleman/strcase"
)

// parseJSONObjects parses the JSON objects and constructs the final struct.
func parseJSONObjects(l *list.List, e map[string]interface{}, tabCount int, isFirst bool) {

	for key, value := range e {
		switch val := value.(type) {
		case string:
			l.PushBack(fmt.Sprintf("%s%-10s%-10s%-10s", printTabs(tabCount), strcase.ToCamel(key), "string", "`json:\""+strcase.ToSnake(key)+"\"`"))
		case []interface{}:
			var tabs int
			if isFirst && tabCount == 1 {
				l.PushBack(fmt.Sprintf("%-10s[]struct %-10s", strcase.ToCamel(key), "{ "))
				isFirst = false
				tabs = tabCount
				tabCount += 1
			}
			t, length := parseJSONArrays(l, val, 1, tabCount, isFirst)
			if len(t) != 0 {
				dataType := getType(t, length)
				l.PushBack(fmt.Sprintf("%s%-10s%-10s%-10s", printTabs(tabCount), strcase.ToCamel(key), dataType, "`json:\""+strcase.ToSnake(key)+"\"`"))
			}
			if tabs == 2 {
				l.PushBack(fmt.Sprintf("%s%-10s%-10s", printTabs(tabs), "}", "`json:\""+strcase.ToSnake(key)+"\"`"))
			}
		case float64:
			l.PushBack(fmt.Sprintf("%s%-10s%-10s%-10s", printTabs(tabCount), strcase.ToCamel(key), "int", "`json:\""+strcase.ToSnake(key)+"\"`"))
		case map[string]interface{}:
			l.PushBack(fmt.Sprintf("%s%-10sstruct %-10s", printTabs(tabCount), strcase.ToCamel(key), "{ "))
			parseJSONObjects(l, val, tabCount+1, isFirst)
			l.PushBack(fmt.Sprintf("%s%s %s", printTabs(tabCount), "}", "`json:\""+strcase.ToSnake(key)+"\"`"))
		}
	}
}

// parseJSONArrays parses the JSON arrys and calls the JSON object parser function for further parsing.
func parseJSONArrays(l *list.List, e []interface{}, level, tabCount int, isFirst bool) ([]string, int) {

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
			dataTypes, level = parseJSONArrays(l, val, level, tabCount, isFirst)
		case map[string]interface{}:
			for key := range val {
				counter[key]++
				lastKey = key
			}
			if counter[lastKey] > 1 {
				continue
			}
			parseJSONObjects(l, val, tabCount, isFirst)
		}
	}
	return dataTypes, level
}
