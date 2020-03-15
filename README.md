# jsonToStruct
A simple application to convert JSON data to go struct.

{
  "name": "Barksalot",
  "species": "Dog",
  "age": 5,
  "photo": "https://learnwebcode.github.io/json-example/images/dog-1.jpg"
}

type auto struct { 
  Name          string	`json:"name"`
  Species       string	`json:"species"`
  Age           int	`json:"age"`
  Photo         string	`json:"photo"`
}

