package lrucache

import (
	lrucacheitem "go-playground/LRU-CacheItem"
	list "go-playground/List"
)

type Lrucache struct {
	limit  int
	list   *list.List
	memory map[string]*lrucacheitem.LRUCacheItem
	size   int
}

func New(limit int) *Lrucache {
	return &Lrucache{limit:limit }
}

func (l *Lrucache) Get(key string) int {
	return l.memory[key].Val
}

func (l *Lrucache) Size() int {
	return l.size
}

func (l *Lrucache) Set(key string, val int) {
	item := lrucacheitem.New(val, key)
	item.Node = l.list.Unshift(item)
	l.memory[key] = item
	l.size++
	if l.limit < l.size {
		oldest := l.list.Pop()
		delete(l.memory, oldest.Key)
		l.size--
	}
}
