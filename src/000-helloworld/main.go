package main

import "fmt"

func main() {
	var i int = 0

	x := 42 + 35

	if x+i > 0 {
		fmt.Printf("Hello World 1\n")
	} else {
		fmt.Printf("Hello World 2\n")
	}

	foo(x)

}

func foo(x int) {
	fmt.Printf("X value: %d", x)
}
