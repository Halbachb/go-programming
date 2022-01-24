package main

import "fmt"

func main() {
	fmt.Println("Hello everyone. Thisi s the most fun you will ever have learning the Go programming langauge")

	foo()
	fmt.Println("something else")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("I am in foo")
}

func bar() {
	fmt.Println("and then we exited control flow")
}

// control flow:
// (1) sequence
// (2) loop; iterative
// (3) conditional
