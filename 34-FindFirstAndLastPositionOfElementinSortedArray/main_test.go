package main

import "testing"

func Test_(t *testing.T) {
	// simple check
	q, a := test1()
	searchRange(q, a)

	q, a = test2()
	searchRange(q, a)

	q, a = test3()
	searchRange(q, a)
}

func test1() (q []int, target int){
	q = []int {5, 7, 7, 8, 8, 10}
	target = 8
	return q, target
}


func test2() (q []int, target int){
	q = []int {5,7,7,8,8,10}
	target = 6
	return q, target
}

func test3() (q []int, target int){
	q = []int {}
	target = 0
	return q, target
}