package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Person Representation of a single person as a struct (object)
type Person struct {
	Name string
	Age  int
}

// People Repersentation of all the people as a struct
type People struct {
	People []Person
}

func main() {
	myPeople := People{}
	// Read the data from the file (returns as array of bytes)
	jsonBytes, err := ioutil.ReadFile("./peoples.json")
	fmt.Println(string(jsonBytes))
	errorLog(err)
	// Unmarshel the JSON bytes into the mem address of myPeople var
	err = json.Unmarshal(jsonBytes, &myPeople)
	errorLog(err)

	fmt.Println(myPeople.People[0])
}

func errorLog(err error) {
	if err != nil {
		panic(err)
	}
}
