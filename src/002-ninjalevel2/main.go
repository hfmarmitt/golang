package main

import "fmt"

const (
	a     = 10
	b int = 12
)

const (
	ano1 = 2021 - iota
	ano2 = 2021 - iota
	ano3 = 2021 - iota
	ano4 = 2021 - iota
)

func main() {
	handsOnOne()
	handsOnTwo()
	handsOnThree()
	handsOnFour()
	handsOnFive()
	handsOnSix()
}

func handsOnOne() {
	x := 255

	fmt.Printf("%d\t%b\t%X\n", x, x, x)
}

func handsOnTwo() {
	g := (10 == 10)
	h := (10 <= 10)
	i := (11 >= 12)
	j := (0 != 1)
	k := (1 < 2)
	l := (2 > 1)

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
}

func handsOnThree() {
	fmt.Println(a, b)
}

func handsOnFour() {
	x := 42
	fmt.Printf("%d\t%b\t%X\n", x, x, x)

	y := x << 1
	fmt.Printf("%d\t%b\t%X\n", y, y, y)
}

func handsOnFive() {
	var str string = `this is a raw
						string
						literal "testing"`
	fmt.Println(str)
}

func handsOnSix() {
	fmt.Printf("%d\t%d\t%d\t%d\n", ano1, ano2, ano3, ano4)
}
