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
### Application in action

![json-to-struct](https://user-images.githubusercontent.com/6292363/78546973-ecc4ec00-781b-11ea-9a74-63060dc9c726.JPG)

![App in action #2](https://user-images.githubusercontent.com/6292363/113766900-a0b4cc00-973b-11eb-93c5-b9ccea1368e2.png)


### Usage
```sh
./jsonToStruct <file1.json> <file2.json> ..
```

### Installation
```sh
cd jsonToStruct
make exe
```

### Dependencies (using go modules)

  - https://github.com/iancoleman/strcase (latest)

### prerequisites
  
  - Golang 1.11
  
### Todos

 - Write Unit Tests
 - Add and test with more JSON examples
