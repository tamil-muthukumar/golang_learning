package main

import (
	"fmt"
	"github.com/sourcegraph/conc"
	"github.com/sourcegraph/conc/pool"
)

func main() {
	// call process function with a stream of integers
	stream := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			stream <- i
		}
		close(stream)
	}()
	process(stream)
}

func process(stream chan int) {

	// just the results
	//p := pool.NewWithResults[int]().WithMaxGoroutines(10)

	// results and errors
	p := pool.NewWithResults[int]().WithErrors().WithMaxGoroutines(10)

	// no results
	//p := pool.New().WithMaxGoroutines(10)

	for elem := range stream {
		elem := elem
		p.Go(func() (int, error) {
			return handleElem(elem)
		})
	}
	results, err := p.Wait()
	fmt.Println("Results: ", results, err)
}

func handleElem(elem int) (int, error) {
	if elem == 5 {
		panic("intentional panic")
	}
	fmt.Println("Handling element: ", elem)
	return elem, nil
}

func simpleWaitGroup() {
	var wg conc.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Go(func() {
			// do something
			fmt.Println("Doing something, index: ", i)
		})
	}
	wg.Wait()
}
