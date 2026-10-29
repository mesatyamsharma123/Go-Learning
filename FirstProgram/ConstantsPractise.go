package main

import "fmt"

func ConstantPractise() {
	const pi = 3.14159
	const name = "Alice"
	fmt.Println(pi)
	fmt.Println(name)

	const (
		a = 2
		b = 3
	)
	fmt.Println(a)
	fmt.Println(b)
}
