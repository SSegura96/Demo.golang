package structure

import "fmt"

func Exercise1() {
	type person struct {
		name     string
		age      int
		location string
	}

	sergio := person{name: "Sergio", age: 25, location: "San Jose, Costa Rica"}
	andrei := person{name: "Andrei", age: 50, location: "bucharest romania"}

	fmt.Printf("#1: %+v\n#2: %+v", sergio, andrei)
}

func Exercise2() {
	type color struct {
		name string
	}

	type person struct {
		name     string
		age      int
		location string
		favColor color
	}

	sergio := person{name: "Sergio", age: 25, location: "San Jose, Costa Rica"}
	sergio.favColor.name = "Blue"

	andrei := person{name: "Andrei", age: 50, location: "bucharest romania"}
	andrei.favColor.name = "Red"

	fmt.Printf("#1: %+v\n#2: %+v", sergio, andrei)
}

func Exercise3() {

	type person struct {
		name      string
		age       int
		location  string
		favColors []string
	}

	me := person{name: "Sergio", age: 25, location: "San Jose, Costa Rica"}
	me.favColors = []string{"Blue", "Orange"}

	fmt.Print("My fav colors: ")
	for _, color := range me.favColors {
		fmt.Print(color, " ")
	}
}

func Exercise4() {
	type grades struct {
		grade  int
		course string
	}

	type person struct {
		name      string
		age       int
		location  string
		gradeList []grades
	}

	sergio := person{name: "Sergio", age: 25, location: "San Jose, Costa Rica", gradeList: []grades{{grade: 98, course: "Programming"}, {grade: 80, course: "Math"}}}
	andrei := person{name: "Andrei", age: 50, location: "bucharest romania", gradeList: []grades{{grade: 40, course: "Programming"}, {grade: 100, course: "Math"}}}

	andrei.gradeList[0].course = "Golang"
	andrei.gradeList[0].grade = 98

	fmt.Printf("#1: %+v\n#2: %+v", sergio, andrei)
}
