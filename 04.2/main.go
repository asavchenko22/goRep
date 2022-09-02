package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func createList(s []int) {
	node := &ListNode{s[0], &ListNode{}}
	for k, v := range s {
		if k > 0 {
			node.next = &ListNode{v, &ListNode{}}
			node = node.next
		}
		fmt.Println(node)
	}

}

func main() {
	s := []int{1, 2, 3}
	createList(s)

}
