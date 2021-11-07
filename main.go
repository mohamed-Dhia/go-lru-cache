package main

import (
	"fmt"
	lrucache "mohamed-Dhia/lru-cache/LRU-Cache"
)

func main() {
	cache := lrucache.New(3)
	cache.Set("item1", 1)
	cache.Set("item2", 2)
	cache.Set("item3", 3)
	cache.Set("item4", 4)
	var item int
	var ok bool
	item, ok = cache.Get("item3")
	fmt.Println(item, ok, 3)
	item, ok = cache.Get("item2")
	fmt.Println(item, ok, 2)
	item, ok = cache.Get("item1")
	fmt.Println(item, ok, nil)
	cache.Set("item5", 5)
	cache.Set("item6", 6)
	item, ok = cache.Get("item3")
	fmt.Println(item, ok, nil)
	item, ok = cache.Get("item4")
	fmt.Println(item, ok, nil)
}
