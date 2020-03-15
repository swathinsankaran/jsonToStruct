package main

import (
	"fmt"
	"reflect"

	"github.com/iancoleman/strcase"
)

// parseJSONObjects parses the JSON objects and constructs the final struct.
func parseJSONObjects(e map[string]interface{}, tabCount int, finalStruct *string) {

	*finalStruct += "\n"
	for key, value := range e {
		*finalStruct += printTabs(tabCount)
		switch val := value.(type) {
		case string:
			*finalStruct += fmt.Sprintf("%-14s%s\t%s", strcase.ToCamel(key), "string", "`json:\""+strcase.ToSnake(key)+"\"`\n")
		case []interface{}:
			var tabs int
			if first && tabCount == 2 {
				*finalStruct += fmt.Sprintf("%s []struct %s", strcase.ToCamel(key), "{ ")
				first = false
				tabs = tabCount
				tabCount += 2
			}
			t, l := parseJSONArrays(val, 1, tabCount, finalStruct)
			if len(t) != 0 {
				dataType := getType(t, l)
				*finalStruct += fmt.Sprintf("%-14s%s\t%s", strcase.ToCamel(key), dataType, "`json:\""+strcase.ToSnake(key)+"\"`\n")
			}
			if tabs == 2 {
				*finalStruct += printTabs(tabs)
				*finalStruct += fmt.Sprintf("%s %s", "}", "`json:\""+strcase.ToSnake(key)+"\"`\n")
			}
		case float64:
			*finalStruct += fmt.Sprintf("%-14s%s\t%s", strcase.ToCamel(key), "int", "`json:\""+strcase.ToSnake(key)+"\"`\n")
		case map[string]interface{}:
			*finalStruct += fmt.Sprintf("%s struct %s", strcase.ToCamel(key), "{ ")
			parseJSONObjects(val, tabCount+2, finalStruct)
			*finalStruct += printTabs(tabCount)
			*finalStruct += fmt.Sprintf("%s %s", "}", "`json:\""+strcase.ToSnake(key)+"\"`\n")
		default:
			fmt.Println(reflect.TypeOf(val))
			continue
		}

	}
}

// parseJSONArrays parses the JSON arrys and calls the JSON object parser function for further parsing.
func parseJSONArrays(e []interface{}, level, tabCount int, finalStruct *string) ([]string, int) {
	if len(e) == 0 {
		panic("Invalid JSON provided")
	}
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
			dataTypes, level = parseJSONArrays(val, level, tabCount, finalStruct)
		case map[string]interface{}:
			for key := range val {
				counter[key]++
				lastKey = key
			}
			if counter[lastKey] > 1 {
				continue
			}
			parseJSONObjects(val, tabCount, finalStruct)
		default:
			fmt.Println(reflect.TypeOf(val))
			continue
		}
	}
	return dataTypes, level
}
