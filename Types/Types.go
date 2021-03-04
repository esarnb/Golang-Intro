package main

import (
	"encoding/json"
	"fmt"
)

// Declare out own custom types
// type Name format
// Formats:
// - struct (objects)
// - interface (advanced level stuff)
// Structs:
/* struct {
	PropName type
} */

// Person has a Name and an Age
type Person struct {
	Name string
	Age  int
}

// MyJSONType JSON type interpected as Go struct
type MyJSONType struct {
	Abc string `json:"fig"`
}

// StringVar with json as string
var strJSON string = "{\"fig\": \"xd\"}"

func main() {
	Emp := Person{
		Name: "Emp",
		Age:  23,
	}
	// Empty MyJSON type
	MyJSON := MyJSONType{}
	err := json.Unmarshal([]byte(strJSON), &MyJSON)
	if err != nil {
		panic(err)
	}
	fmt.Println(Emp.Name, Emp.Age)
	fmt.Println(MyJSON.Abc)

	// Slice:
	// name := []type{val1, val,2}

	// 'any' type:
	// name := []interface{1, "a", 1.5}

	xd := []string{"xd", "xd2"}
	// Array:
	// name := make([]type, len, end)
	// Just use slices damn
	fmt.Println(xd)

	// Append item to array
	xd = append(xd, "xd3")
	fmt.Println(xd)
	// Prepend item to array
	xd = append([]string{"xd0"}, xd...)
	fmt.Println(xd)
}
