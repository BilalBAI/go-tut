package main

import "fmt"

func testIf(x int) {
	if 20 > x {
		fmt.Println("20 > x")
	} else {
		fmt.Println("20 < x")
	}
}

func testLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
func testSwitch(day int) {

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}
}
