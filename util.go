package main

import (
	"fmt"
	"log"
)

// getType checks whether the data types are the same. If same it returns
// data type concatenated with its dimention(2D/3D). If the data type is not same
// then it returns interface{} concatenated with its dimention(2D/3D).
func getType(types []string, level int) string {
	var i int
	var l string = getDiamentionOfSlice(level)
	for i < len(types)-1 {
		if types[i] != types[i+1] {
			return l + "interface{}"
		}
		i++
	}
	return l + types[0]
}

// getDiamentionOfSlice returns the dimention of the Slice based on the level.
// it can be 2D, 3D etc
func getDiamentionOfSlice(level int) string {
	var output string
	for level > 0 {
		output += fmt.Sprintf("[]")
		level--
	}
	return output
}

// printTabs returns the number of tabs to put in front of each line based on the level.
func printTabs(count int) string {
	var i int
	var result string
	for i < count {
		result += "\t"
		i++
	}
	return result
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
