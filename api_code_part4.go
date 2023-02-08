package main

// Channels are a way for Goroutines to communicate with each other and synchronize their execution

// A Goroutine is a lightweight thread of execution in the Go programming language.
//  Goroutines are similar to threads, but they are much cheaper in terms of memory and CPU usage, 
// as they are multiplexed onto a smaller number of actual operating system threads. 
// Goroutines are also much easier to create and manage compared to threads, 
// as they are managed by the Go runtime and are scheduled automatically.

// Goroutines are used to run functions concurrently, so that multiple tasks 
// can be performed at the same time. This makes them a powerful tool for 
// writing concurrent and parallel programs, such as network servers, data pipelines, and more.


import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func downloadFile(url string, ch chan string) {
	//stores response and error as returned by the http request
	response, err := http.Get(url)

	//handling failure error
	if err != nil {
		ch <- fmt.Sprintf("Failed to download %s: %s", url, err)
		return
	}

	//closing the body stream of response 
	defer response.Body.Close()

	//creatiing files where the downloads  shud be saved
	file, err := os.Create(url[strings.LastIndex(url, "/")+1:])

	//handling error
	if err != nil {
		ch <- fmt.Sprintf("Failed to create file for %s: %s", url, err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		//url refered from the channel of download
		ch <- fmt.Sprintf("Failed to save data for %s: %s", url, err)
		return
	}
	//printing finishing statement
	ch <- fmt.Sprintf("Finished downloading %s", url)
}

func main() {
	//example domain has samples of files which can be downloaded and used for technical purpose

// 	Example domains
// As described in RFC 2606 and RFC 6761, a number of domains such as example.com and example.org are maintained for documentation purposes. These domains may be used as illustrative examples in documents without prior coordination with us. They are not available for registration or transfer.

	//lists of urls which is to be downloaded
	urls := []string{
		"https://example.com/file1.html",
		"https://example.com/file2.html",
		"https://example.com/file3.html",
	}

	//defining a channel which will communicate the goroutine process
	ch := make(chan string)

	for _, url := range urls {	
		//calling downloadFile function for downloading file, url and channel is passed as argumnet
		go downloadFile(url, ch)
	}
	
	for i := 0; i < len(urls); i++ {
		//printing the value of channel as used in the downloading goroutine
		fmt.Println(<-ch)
	}
}
