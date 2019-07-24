package twosum

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// Brut force
func BrutForceTwoSum(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[j] == target-nums[i] {
				return []int{i, j}
			}
		}
	}

	return nums
}

// Two-pass Hash Table
func TwoPassHashTableTwoSum(nums []int, target int) []int {
	h := make(map[int]int, 0)
	for j, v := range nums {
		h[v] = j
	}

	for i, v := range nums {
		complement := target - v
		if _, exists := h[complement]; exists && h[complement] != i {
			return []int{i, h[complement]}
		}
	}

	return nums
}

// One-pass Hash Table
func OnePassHashTableTwoSum(nums []int, target int) []int {
	h := make(map[int]int, 0)
	for i, v := range nums {
		complement := target - v
		if _, exists := h[complement]; exists {
			return []int{h[complement], i}
		}

		h[v] = i
	}

	return nums
}
