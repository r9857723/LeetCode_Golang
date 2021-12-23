package main

import "fmt"

func main() {

	q1 := []int {4, 9, 5}
	q2 := []int {9, 4, 9, 8, 4}
	rtn := intersection(q1, q2)
	fmt.Println(rtn)
}

func intersection(nums1 []int, nums2 []int) []int {
	tmp := make(map[int]int)
	var rtn []int
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			if n1 == n2 {
				tmp[n1] = 1
			}
		}
	}

	for i, _ := range tmp {
		rtn = append(rtn, i)
	}
	return rtn
}