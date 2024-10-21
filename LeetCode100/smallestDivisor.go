package leetcode100

import (
	"math"
	"sort"
)

func smallestDivisor(nums []int, threshold int) int {
	if len(nums) == 1 {
		return int(math.Ceil(float64(nums[0]) / float64(threshold)))
	}
	sort.Ints(nums)

	left, right, ans := 1, nums[len(nums)-1], 1

	for left <= right {
		mid := left + (right-left)/2
		if checkSmallestDivisor(nums, threshold, mid) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	return ans
}

func checkSmallestDivisor(nums []int, threshold int, divisor int) bool {
	sum := 0
	for _, num := range nums {
		sum += int(math.Ceil(float64(num) / float64(divisor)))
	}
	return sum <= threshold
}
