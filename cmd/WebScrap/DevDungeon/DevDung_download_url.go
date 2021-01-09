// download_url_DevDungeon.go
// https://www.devdungeon.com/content/web-scraping-go
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter the URL to download:   ")
	myUrl, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	myUrl = strings.TrimSpace(myUrl)

	// Make request
	response, err := http.Get(myUrl)
	if err != nil {
		fmt.Println("\n\n Would you please enter a valid URL ?")
		fmt.Println(" Please rerun the program !\n\n I got this error : ")
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create output file
	fmt.Print("Enter the path/filename.ext to save the download:   ")
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

	// Copy data from HTTP response to file
	_, err = io.Copy(outFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
}
