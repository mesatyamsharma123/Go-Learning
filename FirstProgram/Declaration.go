package main

import "fmt"

func variablePracticeDeclaration() {

	var name string = "satyam"
	fmt.Println(name)

	var age int = 21
	fmt.Println(age)

	var isStudent bool = true
	fmt.Println(isStudent)

	// short variable declaration
	// var a int = 10
	// can be written as
	d := 10
	fmt.Println(d)

	// Group variable declaration
	var (
		a = 10
		b = 20
		c = 30
	)

	fmt.Println(a, b, c)
}
