// Implementing a Sqrt function using Newton's method
//
// Estimate: x := Sqrt(a)
// def f(x) = x^2 - a
// initial guess: v, s.t. f(v) = v^2 - a
// def g(m) = 2v * m + n that tangent to f(x) at point (v,f(v))
// Therefore f(v) = 2v^2 + n = v^2 - a ==> n = -v^2 - a ==> g(m) = 2v*m -v^2 - a
// Let g(v') = 0, ==> 2v*v' -v^2 - a = 0 ==> v' = (v^2 + a) / (2v) where v' is the updated guess
// Update v' = (v^2 + a) / (2v) sufficient times to converge
//
//

package main

import (
	"fmt"
)

func Sqrt(a float64) float64 {
	v := 1.0
	for i := 1; i < 10; i++ {
		v = (v*v + a) / (2 * v)
		fmt.Println(v)
	}
	return v
}
