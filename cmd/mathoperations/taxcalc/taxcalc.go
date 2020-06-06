//Package taxcal will show if you have budget for the total price after inclusing tax
package main

import (
	"fmt"
)

func main() {
	price := 100
	fmt.Println("Price is         ", price, "dollars.")

	taxRate := 0.08
	tax := float64(price) * taxRate
	fmt.Println("Tax   is         ", tax, "dollars.")

	total := float64(price) + tax
	fmt.Println("Total cost is    ", total, "dollars.")

	availableFunds := 120
	fmt.Println("Available Funds  ", availableFunds, "dollars.")
	fmt.Println("Within Budget    ", total <= float64(availableFunds))
}
