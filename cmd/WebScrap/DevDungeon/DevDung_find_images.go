// find_images_in_page.go
// https://www.devdungeon.com/content/web-scraping-go

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	fmt.Print("Enter the URL to fetch link of all images:   ")
	myUrl, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	myUrl = strings.TrimSpace(myUrl)

	// Make HTTP request
	response, err := http.Get(myUrl)
	if err != nil {
		fmt.Println("\n\n Would you please enter a valid URL ?")
		fmt.Println(" Please rerun the program !\n\n I got this error : ")
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find and print image URLs
	document.Find("img").Each(func(index int, element *goquery.Selection) {
		imgSrc, exists := element.Attr("src")
		if exists {
			fmt.Println(imgSrc)
		}
	})
}
