package listnode

type  NodeVal interface {
	
}

type ListNode struct {
	Val  *lrucacheitem.LRUCacheItem
	Prev *ListNode
	Next *ListNode
}

func New(val *lrucacheitem.LRUCacheItem, prev, next *ListNode) *ListNode {
	return &ListNode{ val,  prev,  next}
}

func (l *ListNode) Delete() {
	if l.Prev != nil {
		l.Prev.Next = l.Next
	}
	if l.Next != nil {
		l.Next.Prev = l.Prev
	}
}