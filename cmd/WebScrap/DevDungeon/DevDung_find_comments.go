// find_html_comments_with_regex.go
// https://www.devdungeon.com/content/web-scraping-go
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Print("Enter the URL to fetch all comments:   ")
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

	// Read response data in to memory
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading HTTP body. ", err)
	}

	// Create a regular expression to find comments
	re := regexp.MustCompile("<!--(.|\n)*?-->")
	comments := re.FindAllString(string(body), -1)
	if comments == nil {
		fmt.Println("No matches.")
	} else {
		for _, comment := range comments {
			fmt.Println(comment)
		}
	}
}
