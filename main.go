package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Number of concurrent requests
	concurrentRequests := 50

	// Make multiple requests concurrently
	for i := 0; i < concurrentRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Replace with the URL of your server
			url := "http://localhost:8080"
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading response:", err)
				return
			}

			fmt.Println(string(body))
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
