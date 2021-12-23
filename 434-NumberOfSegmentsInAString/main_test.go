package main

import "testing"

func Test_(t *testing.T) {
	q, a := test1()
	if rtn := countSegments(q); rtn != a {
		t.Errorf("error with test1 = %v , wnat = %v", rtn, a)
	}
	q, a = test2()
	if rtn := countSegments(q); rtn != a {
		t.Errorf("error with test2 = %v , wnat = %v", rtn, a)
	}
}

func test1() (q string, a int) {
	q = "Hello, my name is John"
	a = 5
	return q, a
}
func test2() (q string, a int) {
	q = "Hello"
	a = 1
	return q, a
}