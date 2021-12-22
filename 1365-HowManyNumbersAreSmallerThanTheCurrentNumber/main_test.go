package main

import "testing"

func Test_(t *testing.T){
	// test1
	q, a := test1()
	rtn := smallerNumbersThanCurrent(q)
	fail := false
	for i := 0; i < len(rtn); i++ {
		if rtn[i] != a[i] {
			fail = true
			break
		}
	}
	if fail {
		t.Errorf("error with test1 = %v , wnat = %v", rtn, a)
	}


	// test2
	q, a = test2()
	rtn = smallerNumbersThanCurrent(q)
	fail = false
	for i := 0; i < len(rtn); i++ {
		if rtn[i] != a[i] {
			fail = true
			break
		}
	}
	if fail {
		t.Errorf("error with test2 = %v , wnat = %v", rtn, a)
	}

	// test3
	q, a = test3()
	rtn = smallerNumbersThanCurrent(q)
	fail = false
	for i := 0; i < len(rtn); i++ {
		if rtn[i] != a[i] {
			fail = true
			break
		}
	}
	if fail {
		t.Errorf("error with test3 = %v , wnat = %v", rtn, a)
	}
}

func test1() (q, a []int) {
	q = []int {8,1,2,2,3}
	a = []int {4,0,1,1,3}
	return q, a
}
func test2() (q, a []int) {
	q = []int {6,5,4,8}
	a = []int {2,1,0,3}
	return q, a
}
func test3() (q, a []int) {
	q = []int {7,7,7,7}
	a = []int {0,0,0,0}
	return q, a
}