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
	s := []int{1, 2, 1, 2}
	list := createList(s)
	deleteDuplicates(list)
	fmt.Println(list)
}

func deleteDuplicates(mainNode *ListNode) *ListNode {
	currentNode := mainNode
	var uniqueVal []int
	switchTest := true
	firstValueCheck := true

	for currentNode.next != nil {
		switchTest = true

		if firstValueCheck {
			uniqueVal = []int{currentNode.val}
			firstValueCheck = false
		}

		for k, v := range uniqueVal {
			fmt.Println("v = ", v)
			fmt.Println("currentNode.val = ", v)
			if v != currentNode.val {

				uniqueVal = append(uniqueVal, currentNode.val)
			} else if k > 0 {
				currentNode = currentNode.next
				switchTest = false
			}
		}
		if switchTest {
			currentNode = currentNode.next
		}
	}

	return mainNode
}

//func printList(a *ListNode) *ListNode {
//	return nil
//}
