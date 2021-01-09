// find_links_in_page.go
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

// This will get called for each HTML element found
func processElement(index int, element *goquery.Selection) {
	// See if the href attribute exists on the element
	href, exists := element.Attr("href")
	if exists {
		fmt.Println(href)
	}
}

func main() {
	fmt.Print("Enter the URL :   ")
	myUrl, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	myUrl = strings.TrimSpace(myUrl)

	// Make HTTP request
	response, err := http.Get(myUrl)
	if err != nil {
		fmt.Println("\n\nWould you please enter a valid URL ?")
		fmt.Println("Please rerun the program !\n\n I got this error : ")
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create a goquery document from the HTTP response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find all links and process them with the function
	// defined earlier
	document.Find("a").Each(processElement)
}
