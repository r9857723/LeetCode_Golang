package main

func maximumWealth(accounts [][]int) int {
	rtn := 0
	for _, account := range accounts {
		total := 0
		for _, a := range account {
			total += a
		}
		if rtn < total {
			rtn = total
		}
	}
	return rtn
}