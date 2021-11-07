package lrucache

import (
	lrucacheitem "mohamed-Dhia/lru-cache/LRU-CacheItem"
	list "mohamed-Dhia/lru-cache/List"
)

type Lrucache struct {
	limit  int
	list   *list.List
	memory map[string]*lrucacheitem.LRUCacheItem
	size   int
}

func New(limit int) *Lrucache {
	return &Lrucache{limit: limit, list: list.New(), memory: make(map[string]*lrucacheitem.LRUCacheItem)}
}

func (l *Lrucache) Get(key string) (int, bool) {
	item, ok := l.memory[key]
	if !ok {
		return 0, false
	}
	l.list.MoveToFront(item.Node)
	return item.Val, true
}

func (l *Lrucache) Size() int {
	return l.size
}

func (l *Lrucache) Set(key string, val int) {
	if item, exists := l.memory[key]; exists {
		item.Val = val
		l.list.MoveToFront(item.Node)
	} else {
		item := lrucacheitem.New(val, key)
		item.Node = l.list.Unshift(item)
		l.memory[key] = item
		l.size++
		if l.limit < l.size {
			oldest := l.list.Pop()
			delete(l.memory, oldest.Key())
			l.size--
		}
	}
}
