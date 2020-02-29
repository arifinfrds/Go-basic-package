// Go Packages
// https://golangbot.com/go-packages/

package main

import (
	"fmt"
	"geometry/rectangle"
	"log"
)

var rectLen float64 = 6
var rectWidth float64 = 7

func init() {
	println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero.")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero.")
	}
}

func main() {
	fmt.Println("Geometrical shape properties")
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	fmt.Printf("diagonal of the rectangle %.2f ", rectangle.Diagonal(rectLen, rectWidth))
}
