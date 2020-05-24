package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter the score:   ")
	// reader := bufio.NewReader(os.Stdin)
	// LPAR, err := reader.ReadString('\n')
	score, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	score = strings.TrimSpace(score)

	grade, err := strconv.ParseFloat(score, 64)
	if err != nil {
		fmt.Println("\n\nWould you please enter a number ?")
		fmt.Println("Please rerun the program !\n\n I got this error : ")
		log.Fatal(err)
	}

	if grade == 100 {
		fmt.Println(" Perfect Score !")
	} else if grade >= 40 {
		fmt.Println(" Passed !")
	} else if grade < 40 {
		fmt.Println(" Failed !")
	}
}
