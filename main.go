package main

import (
	"fmt"
	"sync"
	"time"
)

const numIds = 1000

func main() {
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(numIds)
	for i := 0; i < numIds; i++ {
		go func(i int) {
			defer wg.Done()
			u, err := NewID()
			if err != nil {
				panic(err)
			}
			fmt.Printf("UUID %v: %v\n", i, u.String())
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("1000 ids took = %f\n", elapsed.Seconds())
}
