package main

import (
    "net/http"
    "sync"
)

func main() {
    // Create a new HTTP client.
    client := &http.Client{}

    // Create a wait group.
    var wg sync.WaitGroup

    // Generate requests.
    numRequests := 100

    // Create the request URL.
    reqURL := "https://api.example.com/helloweb"

    for i := 0; i < numRequests; i++ {
        // Create a goroutine to make the request.
        wg.Add(1)
        go func() {
            // Make the request.
            client.Get(reqURL)

            // Notify the wait group that the request is complete.
            wg.Done()
        }()
    }

    // Wait for all the requests to complete.
    wg.Wait()
}
