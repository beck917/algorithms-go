package main

import (
	"fmt"
)

func main() {
	fmt.Println(compressString("aabcccccaaa"))

	lru := Constructor(2)

	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1))
	lru.Put(3, 3)
	fmt.Println(lru.Get(2))
}
