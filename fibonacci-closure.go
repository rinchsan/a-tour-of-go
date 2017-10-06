package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
  var secondLast int = 0
	var last int = 1
	return func() int {
		temp := secondLast
		last, secondLast = last + secondLast, last
		return temp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
