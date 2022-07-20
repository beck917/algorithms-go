package main

import (
	"fmt"
	"sync"
)

func main() {
	var ch chan struct{}
	var wg sync.WaitGroup

	ch = make(chan struct{})

	wg.Add(100)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- struct{}{}
			if i%2 == 0 {
				fmt.Println(1, i)
				wg.Done()
			}
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			<-ch
			if i%2 != 0 {
				fmt.Println(2, i)
				wg.Done()
			}
		}
	}()

	wg.Wait()
}
