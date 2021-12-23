package main

import "fmt"

func main() {
	q := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(q))
}
func containsDuplicate(nums []int) bool {
	rtn := make(map[int]bool)
	for _, num := range nums {
		if rtn[num] {
			return true
		}
		rtn[num] = true
	}
	return false
}
