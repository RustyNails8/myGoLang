package main

import (
	"fmt"
	// "math"
	"reflect"
)

func main() {
	// var lenghth float64 = 1.2
	// var width int = 2
	lenghth := 1.2
	// width :=2
	width := int(lenghth)

	fmt.Println("Width is      ", width)
	fmt.Println("Length is     ", lenghth)
	fmt.Println("Type of len  ", reflect.TypeOf(lenghth))
	fmt.Println("Type of Width", reflect.TypeOf(width))
	fmt.Println("Area is      ", lenghth*float64(width))
	fmt.Println("Len > Width ?", lenghth > float64(width))
}
