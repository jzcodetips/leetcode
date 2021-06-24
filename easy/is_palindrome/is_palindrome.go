package main

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	n := 0

	for x > n {
		n = n*10 + x%10
		x = x / 10
	}

	return n == x || x == n/10
}
