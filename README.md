# jsonToStruct

A simple application to convert JSON data to go struct.

Example:

A sample JSON input to the application

```sh
{
  "name": "Barksalot",
  "species": "Dog",
  "age": 5,
  "photo": "https://learnwebcode.github.io/json-example/images/dog-1.jpg"
}
```

gets converted to the following struct.

```sh
type auto struct { 
  Name          string	`json:"name"`
  Species       string	`json:"species"`
  Age           int	`json:"age"`
  Photo         string	`json:"photo"`
}
```
### Usage
```sh
./jsonToStruct <file1.json> <file2.json> ..
```

### Installation
```sh
cd jsonToStruct
make exe
```

### Dependencies

  - https://github.com/iancoleman/strcase

### prerequisites
  
  - Golang 1.11
  
### Todos

 - Write Unit Tests
 - Add and test with more JSON examples
