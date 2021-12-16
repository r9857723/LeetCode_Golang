package main

import (
	"testing"
)

func Test_(t *testing.T) {
	v1, r1 := test1()
	if rtn := maximumWealth(v1); rtn != r1 {
		t.Errorf("error with test1 = %v , wnat = %v", rtn, r1)
	}
	v2, r2 := test2()
	if rtn := maximumWealth(v2); rtn != r2 {
		t.Errorf("error with test1 = %v , wnat = %v", rtn, r2)
	}
	v3, r3 := test3()
	if rtn := maximumWealth(v3); rtn != r3 {
		t.Errorf("error with test1 = %v , wnat = %v", rtn, r3)
	}
}

func test1() ([][]int, int){
	account := [][]int{{1,5},{7,3},{3,5}}
	return account, 10
}

func test2() ([][]int, int){
	account := [][]int{{1,2,3},{3,2,1}}
	return account, 6
}
func test3() ([][]int, int){
	account := [][]int{{2,8,7},{7,1,3},{1,9,5}}
	return account, 17
}