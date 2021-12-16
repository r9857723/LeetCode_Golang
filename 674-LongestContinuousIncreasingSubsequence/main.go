package main

func findLengthOfLCIS(nums []int) int {
	len := len(nums)
	maxCount, count := 1, 1
	for i:= 0; i < len; i++ {
		if i + 1 >= len {
			break
		}
		if nums[i] < nums[i+1] {
			count ++
		} else {
			count = 1
		}
		if maxCount < count {
			maxCount = count
		}
	}
	return maxCount
}