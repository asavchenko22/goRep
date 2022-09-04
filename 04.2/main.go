package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

type Empty struct{}

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
	s := []int{1, 1, 2, 3, 5, 5, 5, 7, 7, 5, 5, 1, 1, 2, 0, 0, 4, 3, 2, 1, 2}
	list := createList(s)
	deleteDuplicates(list)
	fmt.Println(list)
}

func deleteDuplicates(mainNode *ListNode) *ListNode {
	var uniqueVal = map[int]Empty{
		mainNode.val: {},
	}

	currentNode := mainNode
	for currentNode.next != nil {
		if _, ok := uniqueVal[currentNode.next.val]; ok {
			currentNode.next = currentNode.next.next
			continue
		} else {
			uniqueVal[currentNode.next.val] = Empty{}
		}
		currentNode = currentNode.next
	}
	return currentNode
}

//func printList(a *ListNode) *ListNode {
//	return nil
//}
