package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/pkg/profile"
)

// Number of ids we are going to create in this test application
const numIds = 1000

func main() {
	// Enabling a simple profiling in the test application and generating a file to check the memory usage
	defer profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.NoShutdownHook).Stop()

	// Creating a variable with the start time to be able to check how long it took to create a 1000 ids
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(numIds)
	for i := 0; i < numIds; i++ {
		go func(i int) {
			defer wg.Done()
			// Creating the IDs based on RC4122 and adding a timestamp in the front of so it is easier to sort them by creating time.
			u, err := NewIDSortable()
			if err != nil {
				panic(err)
			}
			// Just to  make sure the ids are being created correctly, we wouldn't have this in real life
			fmt.Printf("UUID %v: %v\n", i, u)
		}(i)
	}
	wg.Wait()

	// Printing the time it took to create a 1000 ids.
	elapsed := time.Since(start)
	fmt.Printf("%v ids took = %f\n", numIds, elapsed.Seconds())
}
