package main

import "testing"

func Test(t *testing.T) {

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
	reverseList(l1)
	l7 := &ListNode{
		Val: 2,
		Next: nil,
	}
	l8 := &ListNode{
		Val: 1,
		Next: l7,
	}
	reverseList(l8)

	reverseList(nil)
}