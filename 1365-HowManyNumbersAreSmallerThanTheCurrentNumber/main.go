package main

import "fmt"

func main() {

	q := []int {8,1,2,2,3}
	fmt.Println(smallerNumbersThanCurrent(q))

}
func smallerNumbersThanCurrent(nums []int) []int {

	rtn := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		rtn[i] = process(nums[i], nums)
	}
	return rtn
}

func process(num int, nums []int) int {
	var cnt int
	for i := 0; i < len(nums); i++ {
		if nums[i] < num {
			cnt++
		}
	}
	return cnt
}