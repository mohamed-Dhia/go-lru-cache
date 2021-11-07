package lrucacheitem

import (
	"fmt"
	listnode "mohamed-Dhia/lru-cache/ListNode"
)

type LRUCacheItem struct {
	key  string
	Val  int
	Node *listnode.ListNode
}

func New(val int, key string) *LRUCacheItem {
	return &LRUCacheItem{key, val, nil}
}

func (i *LRUCacheItem) Key() string {
	return i.key
}

func (i *LRUCacheItem) String() string {
	return fmt.Sprintf("{key: ¨%s, Val: %d}", i.key, i.Val)
}
