package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {

	client := &http.Client{}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	var wg sync.WaitGroup

	log.Println("Starting:")
	for i := 0; i < 100; i++ {
		go func(i int) {

			wg.Add(1)
			defer wg.Done()

			log.Printf("Entering goroutine: %d\n", i)
			time.Sleep(10 * time.Microsecond)

			req, err := http.NewRequest("GET", "http://127.0.0.1:8080/", nil)
			if err != nil {
				log.Printf("Goroutine: %d - Error creating request: \n"+err.Error(), i)
				return
			}

			req = req.WithContext(ctx)
			select {
			case <-ctx.Done():
				log.Printf("Goroutine: %d - Context Cancelled, Not doing Request.\n", i)
				return //To not leak coroutine
			default:
				resp, err := client.Do(req)
				if ctx.Err() == context.Canceled {
					log.Printf("Goroutine: %d - Context Deadline Exceeded.", i)
					return
				}
				if err != nil {
					log.Printf("Goroutine: %d - Error: \n"+err.Error(), i)
					return
				}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Printf("Goroutine: %d - Error: \n"+err.Error(), i)
					return
				}
				if string(body) == "7" {
					log.Printf("Goroutine: %d - Found 7. Cancelling Context. \n", i)
					cancel()
					return
				}
				log.Printf("Goroutine: %d - Found %s.\n", i, string(body))
			}
		}(i)
	}
	wg.Wait()
	log.Println("Finished.")
}
