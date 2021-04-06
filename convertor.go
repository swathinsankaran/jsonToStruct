package main

import (
	"container/list"
)

type Convertor interface {
	Convert()
	Print()
}

func newConvertor(convertorType, filePath string) Convertor {
	switch convertorType {
	case "JSONToStruct":
		return &jsonToStruct{
			file:            newFile(filePath),
			generatedStruct: list.New(),
		}
	}
	return nil
}
