package lrucacheitem

import (
	listnode "go-playground/ListNode"
)


type LRUCacheItem struct {
	Key string 
	Val int
	Node *listnode.ListNode
}

func New ( val int,key string) *LRUCacheItem {
	return &LRUCacheItem{key, val, nil }
}