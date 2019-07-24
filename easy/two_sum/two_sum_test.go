package twosum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	arrays := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{2, 8, 20, 17}, 19, []int{0, 3}},
	}

	for _, array := range arrays {
		a := BrutForceTwoSum(array.nums, array.target)
		if a[0] != array.expected[0] || a[1] != array.expected[1] {
			t.Errorf("BrutForceTwoSum got %v, expected %v", a, array.expected)
		}

		b := TwoPassHashTableTwoSum(array.nums, array.target)
		if b[0] != array.expected[0] || b[1] != array.expected[1] {
			t.Errorf("TwoPassHashTableTwoSum got %v, expected %v", b, array.expected)
		}

		c := OnePassHashTableTwoSum(array.nums, array.target)
		if c[0] != array.expected[0] || c[1] != array.expected[1] {
			t.Errorf("OnePassHashTableTwoSum got %v, expected %v", c, array.expected)
		}
	}
}
