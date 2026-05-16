package main

import (
	"fmt"
	"slices"
	"testing"
)

type List struct {
	data int
	next *List
}

func NewList(values ...int) *List {
	root := &List{data: values[0]}
	next := root
	for _, v := range values[1:] {
		next.next = &List{data: v}
		next = next.next
	}
	return root
}

func (list *List) AddNode(value int) {
	var node *List
	for node = list; node.next != nil; node = node.NextNode() {
	}
	node.next = &List{data: value}
}

func (node List) NextNode() *List {
	return node.next
}

func (list List) Traverse() []int {
	elems := []int{}
	for elem := &list; elem.next != nil; elem = elem.next {
		elems = append(elems, elem.data)
	}
	return elems
}

func (l *List) Reverse() *List {
	var prev *List
	curr := l
	for curr != nil {
		// prev, curr, curr.next = curr, curr.next, prev
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev
}

// func (l *List) String() string {
// 	str := "["
// 	for node := l; node != nil; node = node.NextNode() {
// 		str += strconv.Itoa(node.data) + " "
// 	}
// 	return strings.TrimRight(str, " ") + "]"
// }

func reverseLinkedListHelper(t *testing.T, direct, reversed []int) {
	t.Helper()
	t.Logf("reverse linked list of %v", direct)
	list := NewList(direct...)
	if got := list.Reverse(); !slices.Equal(got, reversed) {

	}
}

func TestReverseLinkedList(t *testing.T) {
	t.Run("Count to ten", func(t *testing.T) {})
}

func main() {
	list := NewList(1, 2, 3, 4, 5)
	list.AddNode(6)
	list.AddNode(7)

	fmt.Println(list)

	list = list.Reverse()

	fmt.Println(list)
}
