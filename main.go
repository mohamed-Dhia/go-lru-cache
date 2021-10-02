package main

import (
	"fmt"
	lrucache "go-playground/LRU-Cache"
)

func main() {
	cache := lrucache.New(3)
	cache.Set("item1", 1)
	cache.Set("item2", 2)
	cache.Set("item3", 3)
	cache.Set("item4", 4)

	fmt.Println(cache.Get("item3"), 3)
	fmt.Println(cache.Get("item2"), 2)
	fmt.Println(cache.Get("item1"), nil)
	cache.Set("item5", 5)
	cache.Set("item6", 6)

}
