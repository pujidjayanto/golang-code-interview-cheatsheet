package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Add(v int) {
	newNode := &Node{Value: v}
	if l.Head == nil {
		l.Head = newNode
		return
	}

	// find last node
	// O(n)
	cur := l.Head
	for cur.Next != nil {
		cur = cur.Next
	}

	cur.Next = newNode
}
func (l *LinkedList) Traverse() {
	cur := l.Head
	for cur != nil {
		fmt.Print(cur.Value, " ")
		cur = cur.Next
	}
	fmt.Println()
}

func (l *LinkedList) Sort() {
	l.Head = mergeSort(l.Head)
}

func split(head *Node) (*Node, *Node) {
	if head == nil || head.Next == nil {
		return head, nil
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil

	return head, mid
}

func merge(l1, l2 *Node) *Node {
	dummy := &Node{}
	tail := dummy

	for l1 != nil && l2 != nil {
		if l1.Value < l2.Value {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}

	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}

	return dummy.Next
}

func mergeSort(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	left, right := split(head)
	left = mergeSort(left)
	right = mergeSort(right)
	return merge(left, right)
}

func main() {
	var list LinkedList
	list.Add(2)
	list.Add(10)
	list.Add(5)
	list.Add(7)
	list.Add(9)
	list.Add(4)

	fmt.Print("Unsorted: ")
	list.Traverse()

	list.Sort()

	fmt.Print("Sorted:   ")
	list.Traverse()
}
