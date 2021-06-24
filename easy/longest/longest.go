package main

func max(i, j int) int {
	if i < j {
		return j
	}

	return i
}

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	i, j, ans := 0, 0, 0

	for j < len(s) {
		c := s[j]

		if v, ok := m[c]; ok {
			i = max(v, i)
		}

		ans = max(ans, j-i+1)

		m[c] = j + 1

		j++
	}

	return ans
}
