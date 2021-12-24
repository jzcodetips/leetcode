package main

func rotate(nums []int, k int) {
	l := len(nums)

	if l <= 1 {
		return
	}

	k = k % l

	rotateToRight(nums, k)
}

func rotateToRight(nums []int, k int) {
	reverse(nums)
	reverse(nums[0:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	start := 0
	end := len(nums) - 1

	var tmp int

	for start < end {
		tmp = nums[start]
		nums[start] = nums[end]
		nums[end] = tmp

		start++
		end--
	}
}
