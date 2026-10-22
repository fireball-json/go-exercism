package linkedlist

import (
	"errors"
)

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.

type Node struct {
	Value any
	next  *Node
	prev  *Node
}

type List struct {
	start *Node
	end   *Node
}

func NewList(elements ...any) *List {
	list := &List{}
	for _, v := range elements {
		list.Push(v)
	}
	return list
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v any) {
	node := &Node{Value: v, prev: nil, next: l.start}
	if l.start == nil {
		l.start = node
		l.end = node
	} else {
		l.start.prev = node
		l.start = node
	}
}

func (l *List) Push(v any) {
	node := &Node{Value: v, prev: l.end, next: nil}
	if l.end == nil {
		l.start = node
		l.end = node
	} else {
		l.end.next = node
		l.end = node
	}

}

func (l *List) Shift() (any, error) {
	if l.start == nil {
		return nil, errors.New("empty list")
	} else {
		if l.start == l.end {
			res := l.start.Value
			l.start = nil
			l.end = nil
			return res, nil
		} else {
			l.start.next.prev = nil
			res := l.start.Value
			l.start = l.start.next
			return res, nil
		}
	}
}

func (l *List) Pop() (any, error) {
	if l.end == nil {
		return nil, errors.New("empty list")
	} else {
		if l.start == l.end {
			res := l.end.Value
			l.start = nil
			l.end = nil
			return res, nil
		} else {
			l.end.prev.next = nil
			res := l.end.Value
			l.end = l.end.prev
			return res, nil
		}
	}
}

func (l *List) Reverse() {
	ok := false
	cur := l.start
	for !ok {
		if cur == nil {
			ok = true
			break
		} else {
			temp := cur.next
			cur.next = cur.prev
			cur.prev = temp
			cur = cur.prev
		}
	}
	l.start, l.end = l.end, l.start
}

func (l *List) First() *Node {
	return l.start
}

func (l *List) Last() *Node {
	return l.end
}

func (l *List) Count() int {
	isEnd := false
	cur := l.start
	count := 0
	for !isEnd {
		if cur == nil {
			isEnd = true
			break
		} else {
			cur = cur.next
			count += 1
		}
	}
	return count
}

// Delete removes the first node in a list with a given value.
// Returns true if a node was removed.
func (l *List) Delete(v any) bool {
	ok := false
	cur := l.start
	for !ok {
		if cur == nil {
			ok = true
			break
		} else {
			if cur.Value == v {
				nextNode := cur.next
				prevNode := cur.prev
				if nextNode == nil {
					l.Pop()
				} else if prevNode == nil {
					l.Shift()
				} else if prevNode != nil && nextNode != nil {
					prevNode.next = nextNode
					nextNode.prev = prevNode
				}
				return true
			} else {
				cur = cur.next
			}
		}
	}
	return false
}
