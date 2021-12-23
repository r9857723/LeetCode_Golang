package main

import (
	"fmt"
)

func main() {
	q := []int {1}
	searchRange(q, 1)

}

func searchRange(nums []int, target int) []int {
	rtn := []int {-1, -1}
	for i, n := range nums {
		if n == target {
			rtn[0] = i
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == target {
			rtn[1] = i
			break
		}
	}
	fmt.Println(rtn)
	return rtn
}