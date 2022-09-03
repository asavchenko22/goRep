package main

type ListNode struct {
	val  int
	next *ListNode
}

func createList(s []int) *ListNode {
	first := ListNode{s[0], &ListNode{}}
	prev := &first
	for k, v := range s {
		if k > 0 {
			prev.next = &ListNode{}
			prev.next.val = v
			prev = prev.next
		}
	}
	return &first
}
func main() {
	s := []int{1, 2, 3}
	a := createList(s)
	print(a)
}
