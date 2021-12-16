package main

import "testing"

func Test_(t *testing.T) {
	t1, r1 := test1()
	if rtn := findLengthOfLCIS(t1); rtn != r1 {
		t.Errorf("error with test1 = %v , wnat = %v", rtn, r1)
	}
	t2, r2 := test2()
	if rtn := findLengthOfLCIS(t2); rtn != r2 {
		t.Errorf("error with test2 = %v , wnat = %v", rtn, r2)
	}
}

func test1() ([]int, int) {
	return []int{1,3,5,4,7}, 3
}
func test2() ([]int, int) {
	return []int{2,2,2,2,2}, 1
}