package main

import (
	"fmt"
	"reflect"
)

type People struct { // People is public
	age      int // age is private, if outside package wanna to access -> Age -> public
	name     string
	isMale   bool
	subjects []string
}

type Animal struct {
	People  // extend from People struct
	species string
	phone   int `Maximum can't over 100`
}

func main() {
	// declare struct map[key]data_type {"key": value, ...}
	// constructor Map struct no value:
	peopleNameAgeMap := make(map[string]int)
	peopleNameAgeMap = map[string]int{
		"Alex":    18,
		"Henrry":  20,
		"Alice": 22,
	}

	// m := map[[]int]int{} => error key of map not a slice
	// m := map[[3]int]int{} => not error

	fmt.Println(peopleNameAgeMap)
	fmt.Println(peopleNameAgeMap["Alex"]) // access data of the key name "Alex"
	peopleNameAgeMap["Alex"] = 30         // update value of the key name "Alex"
	peopleNameAgeMap["Thuy Tien"] = 21          // insert new value with key name "Thuy Tien"
	delete(peopleNameAgeMap, "Thuy Tien")       // delete value of the key name "Thuy Tien"
	_, isExist := peopleNameAgeMap["Thuy Tien"] // param 1: value, param 2: bool
	fmt.Println(isExist)                      // false
	fmt.Println(len(peopleNameAgeMap))

	copyMap := peopleNameAgeMap
	copyMap["Alex"] = 25 // same the memory with peopleNameAgeMap

	// struct
	student := People{
		age:      21,
		name:     "Tien",
		isMale:   true,
		subjects: []string{"BackEnd", "FrontEnd", "Design"},
	}

	fmt.Println(student.subjects)

	peoplePlay := People{}

	peoplePlay.age = 25
	peoplePlay.isMale = true
	peoplePlay.name = "Alex"
	peoplePlay.subjects = []string{"Math", "English"}
	fmt.Println(peoplePlay)

	// quick declare
	studentNoName := struct{ name string }{name: "Happy"}
	copyStruct := studentNoName // not same the memory
	copyStruct.name = "Happy Birth New Year"
	// if wanna same the memory -> copyStruct := &studentNoName
	fmt.Println(studentNoName) // Happy
	fmt.Println(copyStruct)    // Happy Birth Day

	// extend
	studentAnimal := Animal{}
	studentAnimal.name = "Dog"
	studentAnimal.isMale = true
	studentAnimal.age = 5
	studentAnimal.subjects = []string{"Cat", "Turtle", "Bird"}
	studentAnimal.species = "Wolf"
	studentAnimal.phone = 101
	fmt.Println(studentAnimal)

	// validate field
	validate := reflect.TypeOf(studentAnimal)
	field, _ := validate.FieldByName("phone") // value, error message
	fmt.Println(field)

}
