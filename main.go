package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	// Check if a file name was passed as an argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file name as an argument")
		return
	}
	// Open the file
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", fileName, err)
		return
	}
	defer file.Close()
	// Read the file contents

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", fileName, err)
		return
	}
	// Split the file contents into URLS
	URLS := strings.Split(string(fileBytes), "\n")

	// Open the file for writing
	outputFile, err := os.OpenFile("reflected_urls.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Printf("Error opening output file: %v\n", err)
		return
	}
	defer outputFile.Close()

	// Iterate over the URLS in the file
	for _, URL := range URLS {
		URL = strings.TrimSpace(URL)
		if URL == "" {
			continue
		}
		// Send a GET request to the target URL
		response, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer response.Body.Close()
		// Read the response body
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Parse the target URL
		parsedURL, err := url.Parse(URL)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Get the query parameters from the parsed URL
		queryParams := parsedURL.Query()

		// Iterate over the query parameters and check if the value is reflected in the response body
		isReflected := false
		for _, values := range queryParams {
			for _, value := range values {
				if strings.Index(string(body), value) != -1 {
					isReflected = true
					break
				}
			}
			if isReflected {
				// Append the URL to the output file
				if _, err := outputFile.WriteString(URL + "\n"); err != nil {
					fmt.Printf("Error writing to output file: %v\n", err)
					continue
				}
				fmt.Println("Reflected: " + URL)
			}
		}
	}
}
