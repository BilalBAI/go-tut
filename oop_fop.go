package main

import (
	"fmt"
)

//////////////////////////////////////////////////////////////////

func myFunction(x int, y int) int {
	return x + y + a1*a2/a3 ^ a4
}
func testRecursion(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testRecursion(x + 1)
}

func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}

// ////////////////////////////////////////////////////////////////
// OOP
type Person struct {
	name string
	age  int
	job  string
}

func (per *Person) changeName(new_name string) {
	per.name = new_name
}

func (per Person) changeName2(new_name string) {
	per.name = new_name
}

func (per Person) showName() {
	fmt.Println((per.name))
}

func TestPerson() {
	per := Person{"Bilal", 100, "king"}
	per.changeName("BB")
	per.showName()
}

//////////////////////////////////////////////////////////////////

type Integer int

func (a Integer) Equal(b Integer) bool {
	return a == b
}

//////////////////////////////////////////////////////////////////

func run() {
	var x Integer
	var y Integer
	x, y = 10, 15
	fmt.Println(x.Equal(y))

}
