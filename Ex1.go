package main

import (
	"fmt"
	"time"
)

func main() {
	numberOfUrls := 1000
	numberOfWokers := 5
	queChannels := make(chan int, numberOfUrls)
	for i := 0; i < numberOfUrls; i++ {
		queChannels <- i
	}

	for i := 0; i < numberOfWokers; i++ {
		go func(name string) {
			for v := range queChannels {
				crawl(name, v)
			}
		}(fmt.Sprintf("%d", i))
	}
	time.Sleep(time.Second * 5)
}

func crawl(name string, v int) {
	time.Sleep(time.Second / 100)
	fmt.Printf("Worker %s is crawling: %d\n", name, v)
}
