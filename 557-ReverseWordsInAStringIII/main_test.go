package main

import "testing"

func Test_(t *testing.T) {
	q1, a1 := test1()
	if rtn := reverseWords(q1); rtn != a1 {
		t.Errorf("error with test1 = %v , wnat = %v", rtn, a1)
	}
	q2, a2 := test1()
	if rtn := reverseWords(q2); rtn != a2 {
		t.Errorf("error with test1 = %v , wnat = %v", rtn, a2)
	}
}

func test1() (q, a string) {
	q = "s'teL ekat edoCteeL tsetnoc"
	a = "Let's take LeetCode contest"
	return q, a
}

func test2()(q, a string) {
	q = "God Ding"
	a = "doG gniD"
	return q, a
}
