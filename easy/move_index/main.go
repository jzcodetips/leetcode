package main

// [1, 2, 3, 4, 5, 6, 7]

func Solution(a []int, k int) {
	reverse(a)
	reverse(a[0:k])
	reverse(a[k:])
}

func reverse(a []int) {
	start := 0
	end := len(a) - 1

	var tmp int

	for start < end {
		tmp = a[end]
		a[end] = a[start]
		a[start] = tmp

		start++
		end--
	}
}
