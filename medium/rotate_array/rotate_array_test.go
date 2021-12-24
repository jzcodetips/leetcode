package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		label  string
		kInput int
		aInput []int
		wanted []int
	}{
		{
			label:  "with simple case",
			kInput: 3,
			aInput: []int{1, 2, 3, 4, 5, 6, 7},
			wanted: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			label:  "with simple case 2",
			kInput: 5,
			aInput: []int{1, 2, 3, 4, 5, 6, 7, 8},
			wanted: []int{4, 5, 6, 7, 8, 1, 2, 3},
		},
		{
			label:  "with k > len(a) and k = 3",
			kInput: 3,
			aInput: []int{1, 2},
			wanted: []int{2, 1},
		},
		{
			label:  "with k > len(a) and k = 5",
			kInput: 5,
			aInput: []int{1, 2},
			wanted: []int{2, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.label, func(t *testing.T) {
			t.Logf("aInput %v", test.aInput)

			rotate(test.aInput, test.kInput)

			t.Logf("wanted %v", test.wanted)

			assert.Equal(t, test.wanted, test.aInput)
		})
	}
}
