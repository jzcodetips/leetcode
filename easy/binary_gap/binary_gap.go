package main

func Solution(n int) int {
	var currentLargest, largest int

	for n > 0 {
		if (n & 1) == 1 {
			break
		}

		n = n >> 1
	}

	for n > 0 {
		if (n & 1) == 0 {
			currentLargest++
		} else {
			if currentLargest > largest {
				largest = currentLargest
			}

			currentLargest = 0
		}

		n = n >> 1
	}

	return largest
}
