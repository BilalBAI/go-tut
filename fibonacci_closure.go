package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	s := 0
	x := 1
	temp := 0
	return func() int {
		temp = x
		x = x + s
		s = temp
		return x
	}
}

func run_fibonacci(limit int) {
	f := fibonacci()
	for i := 0; i < limit; i++ {
		fmt.Println(f())
	}
}
