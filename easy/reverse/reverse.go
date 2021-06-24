package main

import "math"

func reverse(x int) int {
	sign := 1

	if x < 0 {
		sign = -1
		x = x * sign
	}

	n := 0

	for x > 0 {
		n = n*10 + x%10
		x = x / 10
	}

	if n*sign > math.MaxInt32 {
		return 0
	} else if n*sign < math.MinInt32 {
		return 0
	}

	return n * sign
}
