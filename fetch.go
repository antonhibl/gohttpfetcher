package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// for each provided url in input arguments
	for _, url := range os.Args[1:] {
		// send GET request to the server
		response, error := http.Get(url)
		// if an error has occurred
		if error != nil {
			// print fetch status to stderror
			fmt.Fprintf(os.Stderr "fetch: %v\n", error)
			// exit the program safely
			os.Exit(1)
		}
		// read the response using an io reader
		b, error := ioutil.ReadAll(response.Body)
		// close the response body
		response.Body.Close()
		// if an error has occured
		if error != nil {
			// print fetch status to stderror
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, error)
			// exit the program safely
			os.Exit(1)
		}
		// print the response body from the server
		fmt.Printf("%s", b)
	}
}
