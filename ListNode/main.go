package listnode

import "fmt"

type NodeVal interface {
	Key() string
}

type ListNode struct {
	Val  NodeVal
	Prev *ListNode
	Next *ListNode
}

func New(val NodeVal, prev, next *ListNode) *ListNode {
	return &ListNode{val, prev, next}
}

func (l *ListNode) Delete() {
	if l.Prev != nil {
		l.Prev.Next = l.Next
	}
	if l.Next != nil {
		l.Next.Prev = l.Prev
	}
}

func (l *ListNode) String() string {
	return fmt.Sprintf("{Val: %s, Prev: %s, Next: %s}", l.Val, l.Prev, l.Next)
}
