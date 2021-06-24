package main

import "math"

func myAtoi(s string) int {
	var (
		n    int
		sign int = 1
	)

	idx := 0

	for idx < len(s) {
		if s[idx] != ' ' {
			break
		}

		idx++
	}

	if idx >= len(s) {
		return 0
	}

	if s[idx] == '-' {
		sign = -1
		idx++
	} else if s[idx] == '+' {
		idx++
	}

	for idx < len(s) {
		c := s[idx]

		if c >= '0' && c <= '9' {
			if n > math.MaxInt32/10 || (n == math.MaxInt32/10 && int(c)-int('0') > math.MaxInt32%10) {
				if sign == 1 {
					return math.MaxInt32
				} else {
					return math.MinInt32
				}
			}

			n = n*10 + (int(c) - int('0'))
		} else {
			break
		}

		idx++
	}

	return n * sign
}
