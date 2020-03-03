package main

import (
	"reflect"
	"testing"
)

type addTwoNumbersTests struct {
	l1  *ListNode
	l2  *ListNode
	res []int
}

func TestAddTwoNumbers(t *testing.T) {
	for _, test := range getTests() {
		res := getRes(addTwoNumbers(test.l1, test.l2))
		if !reflect.DeepEqual(test.res, res) {
			t.Errorf("error expecting %v got %v", test.res, res)
		}
	}
}

func TestAddTwoNumbersTwo(t *testing.T) {
	for _, test := range getTests() {
		res := getRes(addTwoNumbersTwo(test.l1, test.l2))
		if !reflect.DeepEqual(test.res, res) {
			t.Errorf("error expecting %v got %v", test.res, res)
		}
	}
}

func getTests() []addTwoNumbersTests {
	tests := []addTwoNumbersTests{
		{
			l1:  fillWith([]int{2, 4, 3}),
			l2:  fillWith([]int{5, 6, 4}),
			res: []int{7, 0, 8},
		},
		{
			l1:  fillWith([]int{5}),
			l2:  fillWith([]int{5}),
			res: []int{0, 1},
		},
	}

	return tests
}

func fillWith(a []int) *ListNode {
	l := &ListNode{}
	start := l
	for i, v := range a {
		l.Val = v
		if i < len(a)-1 {
			l.Next = &ListNode{}
			l = l.Next
		}
	}

	l = nil
	return start
}

func getRes(l *ListNode) []int {
	var a []int
	for l != nil {
		a = append(a, l.Val)
		l = l.Next
	}

	return a
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	b.StopTimer()
	test := getTests()[0]
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = addTwoNumbers(test.l1, test.l2)
	}
}

func BenchmarkAddTwoNumbersTwo(b *testing.B) {
	b.StopTimer()
	test := getTests()[0]
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = addTwoNumbersTwo(test.l1, test.l2)
	}
}
