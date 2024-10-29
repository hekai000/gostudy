package nsum

import "sort"

func threeSumSmaller(nums []int, target int) int {
	sort.Ints(nums)
	length := len(nums)
	if length < 3 {
		return 0
	}
	res := 0

	for i := 0; i < length-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, length-1
		for l < r {
			threeSum := nums[i] + nums[l] + nums[r]
			if threeSum < target {
				res += r - l
				l++
			} else {
				r--
			}
		}
	}
	return res
}
