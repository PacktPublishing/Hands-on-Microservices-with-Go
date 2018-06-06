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
	ctx, _ = context.WithTimeout(ctx, 5*time.Millisecond)

	var wg sync.WaitGroup

	log.Println("Starting:")
	for i := 0; i < 50; i++ {

		//We do the sleep when creating the goroutines. Not on the goroutines.
		time.Sleep(10 * time.Microsecond)

		go func(i int) {
			wg.Add(1)
			defer wg.Done()

			log.Printf("Entering goroutine: %d\n", i)

			req, err := http.NewRequest("GET", "http://127.0.0.1:8080/", nil)
			if err != nil {
				log.Printf("Goroutine: %d - Error creating request: \n"+err.Error(), i)
				return
			}
			req = req.WithContext(ctx)
			resp, err := client.Do(req)
			if ctx.Err() == context.DeadlineExceeded {
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
			log.Printf("Goroutine: %d - Found %s.\n", i, string(body))
		}(i)
	}
	wg.Wait()
	log.Println("Finished.")
}

///GOOGLE EXAMPLE OF MULTIPLE REQUESTS SEE WHAT FINDS IT
