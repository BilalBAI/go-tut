package main

import (
	"fmt"
)

var s string = "asd"
var i int = 12
var f float32 = 1.3
var b bool = true
var a1, a2, a3, a4 int = 1, 3, 5, 7

const CC string = "unchangeable and read-only"

func dataStructures() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := []int{4, 5, 6, 7, 8}
	myslice := []int{1, 2, 3}
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
	var c = make(map[string]string)
	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
	fmt.Printf("c\t%v\n", c)

	fmt.Println((arr1))
	fmt.Println((arr2))
	fmt.Println((myslice))

}

func varDeclare() {
	x := 12
	i = 100
	y := 10
	var (
		b1 float32 = 3
		b2 float32 = 1
		b3 float32 = 1.6
	)
	x = int(b1 * b2 * b3)
	fmt.Println("x:", x)
	fmt.Println("myFunc:", myFunction(x, y))
	fmt.Println("s:", s)
	fmt.Println("CC:", CC)
	fmt.Println("i:", i)
	fmt.Println("f:", f)
	fmt.Println("b:", b)
	fmt.Println()

	var i = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)
}
