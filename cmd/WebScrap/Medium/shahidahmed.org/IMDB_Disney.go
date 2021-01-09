// https://medium.com/@shahidahmed.org/programming-in-go-for-web-scraping-aedf937e769d

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	fmt.Print("Enter the URL to scan IMDB List:   ")
	myURL, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	myURL = strings.TrimSpace(myURL)

	// Create output file
	fmt.Print("Enter the path/filename.ext to save the list:   ")
	outFilePath, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	outFilePath = strings.TrimSpace(outFilePath)

	outFile, err := os.Create(outFilePath)
	if err != nil {
		fmt.Println("\n\n Would you please enter a valid path ?")
		fmt.Println(" Please rerun the program !\n\n I got this error : ")
		log.Fatal(err)
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	writer.Write([]string{"Sl.No.", "Movie", "Released", "Certiifcate", "Genre", "RunTime", "Rating", "Votes", "Gross"})

	c := colly.NewCollector()

	// Before making a request, print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	// Callback when colly finds the entry point to the DOM segment having a movie info
	c.OnHTML(`.lister-item-content`, func(e *colly.HTMLElement) {
		//Locate and extract different pieces information about each movie
		number := e.ChildText(".lister-item-index")
		name := e.ChildText(".lister-item-index ~ a")
		year := e.ChildText(".lister-item-year")
		runtime := e.ChildText(".runtime")
		certificate := e.ChildText(".certificate")
		genre := e.ChildText(".genre")
		rating := e.ChildText("[class='ipl-rating-star small'] .ipl-rating-star__rating")
		vote := e.ChildAttr("span[name=nv]", "data-value")
		gross := e.ChildText(".text-muted:contains('Gross') ~ span[name=nv]")

		// Write all scraped pieces of information to output text file
		writer.Write([]string{
			number,
			name,
			year,
			certificate,
			runtime,
			genre,
			rating,
			vote,
			gross,
		})
	})

	// start scraping the page under the given URL
	c.Visit(myURL)
	fmt.Println("End of scraping: ", myURL)
}
