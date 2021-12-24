package main

import "math"

func main() {
	containsNearbyAlmostDuplicate([]int{1, 0, 1, 1}, 1, 2)
}
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {

	for i, num := range nums {
		start := i
		var end int
		if i+k > len(nums)-1 {
			end = len(nums) - 1
		} else {
			end = i + k
		}
		for j := start; j <= end; j++ {
			if i < j && math.Abs(float64(num-nums[j])) <= float64(t) {
				return true
			}
		}
	}
	return false
}
