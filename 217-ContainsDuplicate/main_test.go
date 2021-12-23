package main

import "testing"

func Test_(t *testing.T) {
	q, a := []int {1,2,3,1}, true
	if rtn := containsDuplicate(q); rtn != a {
		t.Errorf("error with test1 = %v , wnat = %v", rtn, a)
	}

	q, a = []int {1,2,3,4}, false
	if rtn := containsDuplicate([]int {1,2,3,4}); rtn != a {
		t.Errorf("error with test2 = %v , wnat = %v", rtn, a)
	}

	q, a = []int {1,1,1,3,3,4,3,2,4,2}, true
	if rtn := containsDuplicate([]int {1,2,3,1}); rtn != a {
		t.Errorf("error with test3 = %v , wnat = %v", rtn, a)
	}
}
