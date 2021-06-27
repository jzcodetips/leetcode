package main

var single1 = map[string]int{
	"V": 5,
	"L": 50,
	"D": 500,
	"M": 1000,
}

var iComb = map[string]int{
	"I":   1,
	"II":  2,
	"III": 3,
	"IV":  4,
	"IX":  9,
}

var xComb = map[string]int{
	"X":  10,
	"XL": 40,
	"XC": 90,
}

var cComb = map[string]int{
	"C":  100,
	"CD": 400,
	"CM": 900,
}

func getIComb(s string, start, end int) (int, int) {
	var (
		v  int
		ok bool
	)

	if start+2 < end {
		v, ok = iComb[s[start:start+3]]
		if ok {
			return start + 3, v
		}
	}

	if start+1 < end {
		v, ok = iComb[s[start:start+2]]
		if ok {
			return start + 2, v
		}
	}

	v = iComb[string(s[start])]

	return start + 1, v
}

func getXComb(s string, start, end int) (int, int) {
	var (
		v  int
		ok bool
	)

	if start+1 < end {
		v, ok = xComb[s[start:start+2]]
		if ok {
			return start + 2, v
		}
	}

	v = xComb[string(s[start])]

	return start + 1, v
}

func getCComb(s string, start, end int) (int, int) {
	var (
		v  int
		ok bool
	)

	if start+1 < end {
		v, ok = cComb[s[start:start+2]]
		if ok {
			return start + 2, v
		}
	}

	v = cComb[string(s[start])]

	return start + 1, v
}

func romanToIntWithMaps(s string) int {
	var (
		n, total int
	)

	end := len(s)

	idx := 0

	for idx < end {
		switch s[idx] {
		case 'I':
			idx, n = getIComb(s, idx, end)
		case 'X':
			idx, n = getXComb(s, idx, end)
		case 'C':
			idx, n = getCComb(s, idx, end)
		default:
			n = single1[string(s[idx])]

			idx++
		}

		total += n
	}

	return total
}

func romanToInt(s string) int {
	single := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var n int

	end := len(s)

	for idx := 0; idx < end; idx++ {
		if s[idx] == 'I' && idx+1 < end && (s[idx+1] == 'V' || s[idx+1] == 'X') {
			n -= 1
		} else if s[idx] == 'X' && idx+1 < end && (s[idx+1] == 'L' || s[idx+1] == 'C') {
			n -= 10
		} else if s[idx] == 'C' && idx+1 < end && (s[idx+1] == 'D' || s[idx+1] == 'M') {
			n -= 100
		} else {
			n += single[s[idx]]
		}
	}

	return n
}
