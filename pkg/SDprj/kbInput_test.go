// Package kbInput takes inputs from user in various formats
package kbInput

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var userPrompt string
var userNumberRaw string
var userNumberInt int
var userNumberFloat float64

// Prompt Function kbInput.Prompt displays the User Prompt Message
func ExamplePrompt(string) {
	fmt.Print("Enter a number")
}

// GetInt Funciton kbInput.GeGetInt takes the integer from User Input
func ExampleGetInt(userNumberRaw string) int {
	userNumberRaw, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	userNumberRaw = strings.TrimSpace(userNumberRaw)

	userNumberInt, err := strconv.Atoi(userNumberRaw)
	if err != nil {
		fmt.Println("\n\nWould you please enter a whole number ?")
		fmt.Println("Please rerun the program !\n\n I got this error : ")
		log.Fatal(err)
	}

	return userNumberInt

}

// GetFloat Funciton kbInput.GeGetInt takes the decimal number from User Input
func ExampleGetFloat(userNumberRaw string) float64 {
	userNumberRaw, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	userNumberRaw = strings.TrimSpace(userNumberRaw)

	userNumberFloat, err := strconv.ParseFloat(userNumberRaw, 64)
	if err != nil {
		fmt.Println("\n\nWould you please enter a decimal point number ?")
		fmt.Println("Please rerun the program !\n\n I got this error : ")
		log.Fatal(err)
	}

	return userNumberFloat

}
