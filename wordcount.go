package main

import (
	"strings" // Needed to use Split
)

func WordCount(s string) map[string]int {
	// var res map[string]int
	res := make(map[string]int)

	ss := strings.Split(s, " ")

	for _, j := range ss {
		_, ok := res[j]
		if ok {

			res[j] = res[j] + 1
		} else {
			res[j] = 1
		}
	}
	return res
}
