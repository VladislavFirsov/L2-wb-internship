package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-mergeChannels(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("Done after %v\n", time.Since(start))

}

func mergeChannels(channels ...<-chan interface{}) <-chan interface{} {
	mergedChan := make(chan interface{})
	for _, channel := range channels {
		wg := &sync.WaitGroup{}
		wg.Add(len(channels))
		go func(ch <-chan interface{}, wg *sync.WaitGroup) {
			defer wg.Done()
			for data := range ch {
				mergedChan <- data
			}
		}(channel, wg)
		wg.Wait()
		close(mergedChan)
	}
	return mergedChan
}
