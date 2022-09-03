package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func createList(s []int) *ListNode {
	first := ListNode{val: s[0]}
	current := &first
	for k, v := range s {
		if k > 0 {
			current.next = &ListNode{val: v}
			current = current.next
		}
	}
	return &first
}

func main() {
	s := []int{1, 2, 1, 3, 4, 3, 3, 4}
	list := createList(s)
	deleteDuplicates(list)
	fmt.Println(list)
}

func deleteDuplicates(mainNode *ListNode) *ListNode {
	current := mainNode
	for current.next != nil {
		current = current.next
	}

	return mainNode
}

//func printList(a *ListNode) *ListNode {
//	return nil
//}
