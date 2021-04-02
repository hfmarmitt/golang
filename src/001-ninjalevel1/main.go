package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true
var underlyingTypeVar int

func main() {
	handsOnOne()
	handsOnTwo()
	handsOnThree()
	handsOnFour()
	handsOnFive()
}

func handsOnOne() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Printf("%d\n", x)
	fmt.Printf(y + "\n")
	fmt.Println(z)

}

func handsOnTwo() {
	fmt.Println(x, y, z)
	fmt.Printf("%d\n", x)
	fmt.Printf(y + "\n")
	fmt.Println(z)
}

func handsOnThree() {
	s := fmt.Sprintf("%d - %s - %t\n", x, y, z)
	fmt.Print(s)
}

func handsOnFour() {
	type customType int

	var x customType

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}

func handsOnFive() {
	type customType int

	var x customType

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	underlyingTypeVar = int(x)
	fmt.Println(underlyingTypeVar)
}
