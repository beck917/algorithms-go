package main

import (
	"fmt"
	"sync"
)

func main() {
	//fmt.Println("123")
	c1 := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 1; i <= 10; i++ {
			c1 <- 1
			if i%2 == 1 {
				fmt.Print(i)
			}
		}
		wg.Done()
	}()
	go func() {
		for i := 1; i <= 10; i++ {
			<-c1
			if i%2 == 0 {
				fmt.Print(i)
			}
		}
		wg.Done()
	}()
	//c1 <- 1
	wg.Wait()
}

func main1() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c1 := make(chan struct{}, 1)
	c2 := make(chan struct{}, 1)
	done := make(chan struct{})

	go func() {
		for _, v := range nums {
			c2 <- struct{}{}
			if v%2 == 1 {
				fmt.Println(v)
			}
			<-c1
		}
	}()

	go func() {
		for _, v := range nums {
			c1 <- struct{}{}
			if v%2 == 0 {
				fmt.Println(v)
			}
			<-c2
		}
		done <- struct{}{}
	}()

	<-done
}

func print(v int, flag int, c1 chan struct{}, c2 chan struct{}) {
	c2 <- struct{}{}
	if v%2 == flag {
		fmt.Println(v)
	}
	<-c1
}
