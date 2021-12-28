package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	l4 := &ListNode{
		Val:  4,
		Next: l5,
	}
	l3 := &ListNode{
		Val:  3,
		Next: l4,
	}
	l2 := &ListNode{
		Val:  2,
		Next: l3,
	}
	l1 := &ListNode{
		Val:  1,
		Next: l2,
	}
	rtn := reverseList(l1)
	for  {
		fmt.Println(rtn)
		if rtn.Next == nil {
			break
		}
		rtn = rtn.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var preNode *ListNode

	for head != nil {
		next := head.Next
		head.Next = preNode
		preNode = head
		head = next
	}
	tmp := preNode
	for  {
		fmt.Println(tmp)
		if tmp.Next == nil {
			break
		}
		tmp = tmp.Next
	}
	return preNode
}
