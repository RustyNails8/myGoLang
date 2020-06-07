//Package hello is first progeam from "Official - how to code in Go"
package main

import (
	"fmt"

	"example.com/in10c2/hello/morestrings"
	// "github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	// fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
