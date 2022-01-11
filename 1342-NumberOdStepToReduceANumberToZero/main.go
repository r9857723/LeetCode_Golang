package main

import "log"

func main() {
	numberOfSteps(14)
}
func numberOfSteps(num int) int {
	cnt := 0
	for num > 0 {
		cnt ++
		switch num % 2 {
		case 1:
			num -=1
		default:
			num = num/2
		}
	}
	log.Println(cnt)
	return cnt
}