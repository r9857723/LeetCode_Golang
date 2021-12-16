package main

import "testing"

func Test_getDecimalValue(t *testing.T) {
	if test1() != 18880 {
		t.Errorf("error test1.")
	}
	if test2() != 5 {
		t.Errorf("error test2... ")
	}
}

func test1() int {
	l15 := &ListNode{ Val: 0, Next: nil }
	l14 := &ListNode{ Val: 0, Next: l15 }
	l13 := &ListNode{ Val: 0, Next: l14 }
	l12 := &ListNode{ Val: 0, Next: l13 }
	l11 := &ListNode{ Val: 0, Next: l12 }
	l10 := &ListNode{ Val: 0, Next: l11 }
	l9 := &ListNode{ Val: 1,  Next: l10 }
	l8 := &ListNode{ Val: 1,  Next: l9 }
	l7 := &ListNode{ Val: 1,  Next: l8 }
	l6 := &ListNode{ Val: 0,  Next: l7 }
	l5 := &ListNode{ Val: 0,  Next: l6 }
	l4 := &ListNode{ Val: 1,  Next: l5 }
	l3 := &ListNode{ Val: 0,  Next: l4 }
	l2 := &ListNode{ Val: 0,  Next: l3 }
	l1 := &ListNode{ Val: 1,  Next: l2 }
	return getDecimalValue(l1)
}

func test2() int {
	l3 := &ListNode{ Val: 1, Next: nil }
	l2 := &ListNode{ Val: 0, Next: l3 }
	l1 := &ListNode{ Val: 1, Next: l2 }
	return getDecimalValue(l1)

}