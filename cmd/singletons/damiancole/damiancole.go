//Package damiancole is to find how many sheets he has ordered
package main

import (
	"fmt"
)

func main() {
	var qty int
	var length, width float64
	var customerName string

	qty = 4
	length, width = 1.2, 2.4
	customerName = "Damian Cole"

	fmt.Println(customerName)
	fmt.Println("has ordered", qty, "sheets")
	fmt.Println("each with a size of", length, "x", width)
	fmt.Println("and having area", length*width, "sq m")

}
