//Package mathstring to check floor function
package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// var num = math.Floor(0.98)
	// var strng = strings.Title("Check math fucniton floor", num)
	fmt.Println(strings.Title("Check math functionn \nfloor"), math.Floor(0.098))
	fmt.Println(strings.Title("Check math functionn \tfloor"), math.Floor(0.098))
	fmt.Println(strings.Title("Check math functionn \\floor"), math.Floor(0.098))
	fmt.Println(strings.Title("Check math functionn \"floor"), math.Floor(0.098))
}
