package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	// target := rand.Intn(100) + 1
	rand.Seed(seconds)
	target := rand.Intn(100) + 1

	fmt.Println(" \n\n *** Guess The Number Game ***")
	fmt.Print(" Guess a number between 1 to 100:  ")
	fmt.Println(" The target was ... ", target)

	for tries := 0; tries < 10; tries++ {

		fmt.Println("You have", 10-tries, "tries left...")
		guessInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		guessInput = strings.TrimSpace(guessInput)
		guess, err := strconv.Atoi(guessInput)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println(" Too low  a guess ...")
		} else if guess > target {
			fmt.Println(" Too high  a guess ...")
		} else if guess == target {
			fmt.Println(" Correct    guess ...")
			break
		}
	}
	fmt.Println(" The target was ... ", target)
}
