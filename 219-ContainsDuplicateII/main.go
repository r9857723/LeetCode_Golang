package main

func main() {

}

func containsNearbyDuplicate(nums []int, k int) bool {
	tmp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if val, ok := tmp[n]; ok {
			if i - val <= k {
				return true
			}
		}
		tmp[n] = i
	}
	return false
}
