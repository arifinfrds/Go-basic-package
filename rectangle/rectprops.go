// Go Packages
// https://golangbot.com/go-packages/

//rectprops.go
package rectangle

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("rectangle package initialized")
}

func Area(len, wid float64) float64 {
	var area = len * wid
	return area
}

func Diagonal(len, wid float64) float64 {
	var diagonal = math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
