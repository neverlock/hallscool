package main

import (
	"fmt"
	"math/rand"
	"net/http"
//	"net/url"
	"sync"
	//"strconv"
	"bytes"
	//"strconv"
	"mime/multipart"
)

/*
func main() {

	// Create a wait group
	var wg sync.WaitGroup
	Max := 100000
	CONCURRENT := 5

	// Create a channel for sending requests
	requests := make(chan *http.Request, CONCURRENT)

	// Create a counter
	count := 0

	postUrl := "https://pala42.com/ee/post.php"
	// Edit Max your post

	values := make(url.Values)

	// Start two goroutines to send requests
	for i := 0; i < CONCURRENT; i++ {
		go func() {
			for req := range requests {
				user := randomString(10)
				//email := fmt.Sprintf("%s@%s.com", user, domain)
				email := fmt.Sprintf("%s@gmail.com", user)
				pass := randomString(10)

				values.Set("email", email)
				values.Set("pass", pass)
				//input digit like Tel No. use strconv.Itoa(rand.Intn(10000000000))
				//values.Set("entry.1007776359", strconv.Itoa(rand.Intn(10000000000)))
				// Submit form
				resp, err := http.PostForm(postUrl, values)
				if err != nil {
					fmt.Println(err)
					continue
				}

				// Check the response status code
				if resp.StatusCode != http.StatusOK {
					fmt.Println("Error:", resp.Status)
					continue
				}
				// Increment the counter
				count++
				// Print the count in the request
				fmt.Println("Count in the request:", count)

			}

			// Close the channel
			close(requests)
		}()
	}

	// Send 10 requests
	for i := 0; i < Max; i++ {
		// Send the request
		requests <- req
		wg.Add(1)
	}

	// Wait for all requests to finish
	wg.Wait()

}

*/

func main() {
	// URL to send POST request to
	url := "https://pala42.com/ee/post.php"

	// Number of requests to send
	numRequests := 1000

	// Maximum number of concurrent requests
	maxConcurrent := 5

	// Create a buffered channel to control the concurrency
	concurrency := make(chan struct{}, maxConcurrent)

	// Create a wait group to wait for all requests to complete
	var wg sync.WaitGroup

	// Loop to send multiple POST requests concurrently
	for i := 0; i < numRequests; i++ {
		// Add a new request to the wait group
		wg.Add(1)

		// Send request concurrently using Goroutine and channel
		go func(counter int) {
			// Acquire a token from the concurrency channel
			concurrency <- struct{}{}

			// Create a new buffer to store the request body
			body := &bytes.Buffer{}

			// Create a new multipart writer
			writer := multipart.NewWriter(body)

			// Add username and password as form fields
			user := randomString(10)
			//email := fmt.Sprintf("%s@%s.com", user, domain)
			email := fmt.Sprintf("%s@gmail.com", user)
			pass := randomString(10)

			writer.WriteField("email", email)
			writer.WriteField("pass", pass)

			// Close the multipart writer to finalize the request body
			writer.Close()

			// Create a HTTP request with POST method and request body
			req, err := http.NewRequest("POST", url, body)
			if err != nil {
				fmt.Println("Failed to create HTTP request:", err)
				return
			}

			// Set Content-Type header to indicate multipart/form-data
			req.Header.Set("Content-Type", writer.FormDataContentType())

			// Create a HTTP client
			client := &http.Client{}

			// Send the HTTP request
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Failed to send HTTP request:", err)
				return
			}

			defer resp.Body.Close()

			// Check the response status code
			if resp.StatusCode != http.StatusOK {
				fmt.Println("Failed to receive successful response:", resp.Status)
				return
			}

			fmt.Println("POST request sent successfully! Counter:", counter)

			// Release the token back to the concurrency channel
			<-concurrency

			// Mark the request as completed in the wait group
			wg.Done()
		}(i + 1) // Pass the counter value as an argument to the Goroutine
	}

	// Wait for all requests to complete
	wg.Wait()
}

func randomString(l int) string {
	var char string
	buf := make([]byte, l)
	char = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	for i := 0; i < l; i++ {
		buf[i] = char[rand.Intn(len(char)-1)]
	}
	return string(buf)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
