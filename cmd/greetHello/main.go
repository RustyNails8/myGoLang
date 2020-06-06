// Medium https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318
package main

import (
	"fmt"
	// "github.com/logrusorgru/aurora"
)

func main() {
	greetMessage := hello("")
	// fmt.Println(aurora.Yellow(greetMessage))
	fmt.Println(greetMessage)

	greetMessage = hello("Sumit")
	// fmt.Println(aurora.Yellow(greetMessage))
	fmt.Println(greetMessage)

}
