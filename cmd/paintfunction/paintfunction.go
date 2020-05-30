package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func sayHi() {
	// Comment
	fmt.Println("Hi Welcome to Paint function....")
}

func calcArea(wd float64, ht float64) float64 {
	// fmt.Printf("%.2f Liters of Paint required", (wd*ht)/10.0)
	return (wd * ht) / 10.0
}

func wall() float64 {
	fmt.Print("Enter the width of wall..... : ")
	wd, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	wd = strings.TrimSpace(wd)
	width, err := strconv.ParseFloat(wd, 64)
	if err != nil {
		fmt.Println("\n\nWould you please enter a number ?")
		fmt.Println("Please rerun the program !\n\n I got this error : ")
		log.Fatal(err)
	}

	fmt.Print("Enter the height of wall.... : ")
	ht, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	ht = strings.TrimSpace(ht)
	height, err := strconv.ParseFloat(ht, 64)
	if err != nil {
		fmt.Println("\n\nWould you please enter a number ?")
		fmt.Println("Please rerun the program !\n\n I got this error : ")
		log.Fatal(err)
	}
	// var ar float64
	fmt.Printf("%.2f Liters of Paint required for this wall\n", calcArea(width, height))
	fmt.Println("---------------")

	// area := calcArea(width, height)
	return calcArea(width, height)
}

func main() {
	sayHi()
	fmt.Print("How many walls do you have ? .... : ")
	numWall, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	numWall = strings.TrimSpace(numWall)
	wallNumber, err := strconv.Atoi(numWall)
	if err != nil {
		fmt.Println("\n\nWould you please enter a number ?")
		fmt.Println("Please rerun the program !\n\n I got this error : ")
		log.Fatal(err)
	}

	var totalPaint float64
	var paintPerWall float64
	fmt.Println("---------------")
	for aWall := wallNumber; aWall > 0; aWall-- {
		fmt.Printf("Now gathering dimentions of wall # %v \n", aWall)
		paintPerWall = wall()
		totalPaint += paintPerWall
	}
	fmt.Println("##############")
	fmt.Printf("Total Liters Paint required .... %.2f", totalPaint)

}
