package main

import (
	"fmt"
	"sync"
)

type empty interface{}
type semaphore chan empty

// acquire n resources
func (s semaphore) P(n int) {
	e := empty{}
	for i := 0; i < n; i++ {
		s <- e
	}
}

// release n resources
func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

/* mutexes */

func (s semaphore) Lock() {
	s.P(1)
}

func (s semaphore) Unlock() {
	s.V(1)
}

/* signal-wait */

func (s semaphore) Signal() {
	s.V(1)
}

func (s semaphore) Wait(n int) {
	s.P(n)
}

const numIds = 1000

func main() {
	var wg sync.WaitGroup

	wg.Add(numIds)
	sem := make(semaphore, numIds)
	for i := 0; i < numIds; i++ {
		go func(i int) {
			defer wg.Done()
			u, err := NewID()
			if err != nil {
				panic(err)
			}
			fmt.Printf("UUID %v: %v\n", i, u.String())
			sem.Signal()
		}(i)
	}
	sem.Wait(numIds)
	wg.Wait()
}
