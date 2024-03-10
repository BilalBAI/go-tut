package main

import (
	"math"
)

// Methods and Functions are similar in nature
// However, methods with value or pointer receivers can
// take either a value or a pointer as the receiver when they are called
// functions that take a value argument must take a value of that specific type
// functions with a pointer argument must take a pointer

type Vertex struct {
	X, Y float64
}

// Methods with pointer receiver
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Method with value receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func testmethods() {
	v1 := Vertex{3, 4} // v1 is a value
	v1.Scale(10)
	v1.Abs()

	v2 := &Vertex{4, 3}
	v2.Scale(10)
	v2.Abs()
}
