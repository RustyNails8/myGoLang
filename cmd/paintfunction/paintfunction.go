package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sayHi() {
	// Comment
	fmt.Println("Hi Welcome to Paint function....")
}

func calcArea(wd float64, ht float64) {
	fmt.Printf("%.2f Liters of Paint required", (wd*ht)/10.0)
}

func main() {
	sayHi()
	fmt.Print("Enter the width of wall..... : ")
	wd, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	wd = strings.TrimSpace(wd)
	width, _ := strconv.ParseFloat(wd, 64)
	fmt.Print("Enter the height of wall.... : ")
	ht, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	ht = strings.TrimSpace(ht)
	height, _ := strconv.ParseFloat(ht, 64)

	// var ar float64
	// fmt.Printf("%.2f Liters of Paint required", ar/10.0)

	// area := calcArea(width, height)
	calcArea(width, height)

}
