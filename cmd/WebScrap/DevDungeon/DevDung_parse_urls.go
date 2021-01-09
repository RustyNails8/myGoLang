// parse_myUrls.go
// https://www.devdungeon.com/content/web-scraping-go
package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	// Parse a complex URL
	// complexUrl := "https://www.example.com/path/to/?query=123&this=that#fragment"
	fmt.Print("Enter the URL :   ")
	myUrl, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	myUrl = strings.TrimSpace(myUrl)

	// complexUrl := "https://www.flipkart.com/search?q=silicone+keyboard&sid=6bo%2Cai3%2C3oe&as=on&as-show=on&otracker=AS_QueryStore_OrganicAutoSuggest_4_17_na_na_na&otracker1=AS_QueryStore_OrganicAutoSuggest_4_17_na_na_na&as-pos=4&as-type=RECENT&suggestionId=silicone+keyboard%7CKeyboards&requestId=d261c9bd-4030-4078-b3c3-983ce594e4d9&as-backfill=on"
	parsedUrl, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("\n\nWould you please enter a valid URL ?")
		fmt.Println("Please rerun the program !\n\n I got this error : ")
		log.Fatal(err)
	}

	// Print out URL pieces
	fmt.Println("Scheme: " + parsedUrl.Scheme)
	fmt.Println("Host: " + parsedUrl.Host)
	fmt.Println("Path: " + parsedUrl.Path)
	fmt.Println("Query string: " + parsedUrl.RawQuery)
	fmt.Println("Fragment: " + parsedUrl.Fragment)

	// Get the query key/values as a map
	fmt.Println("\nQuery values:")
	queryMap := parsedUrl.Query()
	fmt.Println(queryMap)

	// Craft a new URL from scratch
	var customURL url.URL
	customURL.Scheme = "https"
	customURL.Host = "google.com"
	newQueryValues := customURL.Query()
	newQueryValues.Set("Site", "Flipkart")
	newQueryValues.Set("Qtree", "Search")
	customURL.Fragment = "bookmarkLink"
	customURL.RawQuery = newQueryValues.Encode()

	fmt.Println("\nCustom URL:")
	fmt.Println(customURL.String())
}
