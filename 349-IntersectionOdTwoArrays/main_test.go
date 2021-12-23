package main

import "testing"

func Test_(t *testing.T) {
	q1, q2, a := test1()
	if rtn := len(intersection(q1, q2)); rtn != len(a) {
		// simple check
		t.Errorf("error with test1 = %v , wnat = %v", rtn, a)
	}
	q1, q2, a = test2()
	if rtn := len(intersection(q1, q2)); rtn != len(a) {
		// simple check
		t.Errorf("error with test2 = %v , wnat = %v", rtn, a)
	}
}

/*
Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
*/

func test1() (q1, q2, a []int) {
	q1 = []int {1, 2, 2, 1}
	q2 = []int {2, 2}
	a = []int {2}
	return q1, q2, a
}
/*
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Explanation: [4,9] is also accepted.
*/
func test2() (q1, q2, a []int) {
	q1 = []int {4, 9, 5}
	q2 = []int {9, 4, 9, 8, 4}
	a = []int {9, 4}
	return q1, q2, a
}