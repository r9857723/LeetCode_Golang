package twosum

func twoSum(nums []int, target int) []int {
	var record = map[int]int{}
	for i, v := range nums {
		diff := target - v
		if d, ok := record[diff]; ok {
			return []int{d, i}
		}
		record[v] = i
	}
	return []int{}
}
