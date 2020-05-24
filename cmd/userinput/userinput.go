package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter the server name:   ")
	// reader := bufio.NewReader(os.Stdin)
	// LPAR, err := reader.ReadString('\n')
	LPAR, err := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Print(LPAR, "RC:", err)
}
