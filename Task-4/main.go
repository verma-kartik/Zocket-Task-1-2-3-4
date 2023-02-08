package main

import (
	"fmt"
	"net/http"
	"time"
)

var urls = []string{
	"https://www.rubyconf.org/",
	"https://golang.org/",
	"https://www.microsoft.com/en-in",
	"https://kubernetes.io/",
}

type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

func asyncHttpGets(urls []string) []*HttpResponse {
	ch := make(chan *HttpResponse)
	var responses []*HttpResponse
	client := http.Client{}
	for _, url := range urls {

		// Function that takes a string argument representing an url
		// and prints the string and uses net/http to fetch the web resource

		go func(url string) {
			fmt.Printf("Fetching %s \n", url)
			resp, err := client.Get(url)

			// create instance of HttpResponseType & sent to channel
			ch <- &HttpResponse{url, resp, err}
			if err != nil && resp != nil && resp.StatusCode == http.StatusOK {
				resp.Body.Close()
			}
		}(url)
	}

	//checks if something is in the channel
	for {
		select {
		case r := <-ch:
			fmt.Printf("%s was fetched\n", r.url)
			if r.err != nil {
				fmt.Println("with an error", r.err)
			}
			//if there is something, allocate the data to r
			responses = append(responses, r)
			if len(responses) == len(urls) {
				return responses
			}
		case <-time.After(50 * time.Millisecond):
			fmt.Printf(".")
		}
	}
	return responses
}

func main() {
	results := asyncHttpGets(urls)
	for _, result := range results {
		if result != nil && result.response != nil {
			fmt.Printf("%s result.response.Status")
		}
	}
}
