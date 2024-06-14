package main

import (
	"fmt"
	"time"
)

func joinTwoStrings(prefix string, suffix string) string {
	return prefix + suffix
}

func printOutDate() {
	t := time.Now()
	fmt.Println(t)
}

func waitForIt(message string) {
	defer fmt.Println("Done!")
	fmt.Println("Waiting")
	fmt.Println("Waiting")
	fmt.Println("Waiting")
	fmt.Println(message)
}

func areaOfSquare(l, b int64) (int64, int64) {
	var area int64
	var perimeter int64
	area = l * b
	perimeter = l * 4
	return area, perimeter
}

func main() {
	printOutDate()
	waitForIt(joinTwoStrings("Hi", " there"))
	var Length int64 = 5
	var Breadth int64 = 5
	var area, perimeter int64
	area, perimeter = areaOfSquare(Length, Breadth)

	fmt.Println("The area of the square is", area, "The perimeter of the square is", perimeter)

}
