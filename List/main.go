package list

import listnode "go-playground/ListNode"

type List struct {
	Head *listnode.ListNode
	Tail *listnode.ListNode
}

func New ()*List {
	return &List{}
}

func (l *List) Unshift(val interface{}) *listnode.ListNode {
	if l.Head == nil && l.Tail == nil {
		l.Head = listnode.New(val, nil, nil)
		l.Tail = l.Head
	} else {
		l.Head = listnode.New(val, nil, l.Head)
		l.Head.Next.Prev = l.Head
	}
	return l.Head
}

func (l *List) Shift() interface{} {
	if l.Head == nil && l.Tail == nil {
		return nil
	} else {
		h := l.Head
		l.Head = l.Head.Next
		h.Delete()
		return h.Val
	}
}

func (l *List) Push(val interface{}) *listnode.ListNode {
	if l.Head == nil && l.Tail == nil {
		l.Head = listnode.New(val, nil, nil)
		l.Tail = l.Head
	} else {
		l.Tail = listnode.New(val, l.Tail, nil)
		l.Tail.Prev.Next = l.Tail
	}
	return l.Tail
}

func (l *List) Pop() interface{} {
	if l.Head == nil && l.Tail == nil {
		return nil
	} else {
		t := l.Tail
		l.Tail = l.Tail.Prev
		t.Delete()
		return t.Val
	}
}

func (l *List) MoveToFront(node *listnode.ListNode) {
	if node == l.Tail {
		l.Pop()
	} else if node == l.Head {
		return
	} else {
		node.Delete()
	}

	node.Prev = nil
	node.Next = nil

	if l.Head == nil && l.Tail == nil {
		l.Head = node
		l.Tail = l.Head
	} else {
		l.Head.Prev = node
		node.Next = l.Head
		l.Head = node
	}
}

func (l *List) MoveToEnd(node *listnode.ListNode) {
	if node == l.Head {
		l.Shift()
	} else if node == l.Tail {
		return
	} else {
		node.Delete()
	}

	node.Prev = nil
	node.Next = nil

	if l.Head == nil && l.Tail == nil {
		l.Head = node
		l.Tail = l.Head
	} else {
		l.Tail.Next = node
		node.Prev = l.Tail
		l.Tail = node
	}

}
