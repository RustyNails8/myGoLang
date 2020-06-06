// Medium https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318
package main

import (
	"fmt"
)

// hello to print hello
func hello(user string) string {
	if len(user) == 0 {
		return fmt.Sprintf("hello Dude !!")
	} else {
		return fmt.Sprintf("hello %v !!", user)
	}

}
