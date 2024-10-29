package nsum

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	length := len(nums)

	res := nums[0] + nums[1] + nums[length-1]
	for i := 0; i < length-2; i++ {
		l, r := i+1, length-1
		for l < r {
			threeSum := nums[i] + nums[l] + nums[r]
			if math.Abs(float64(threeSum-target)) < math.Abs(float64(res-target)) {
				res = threeSum
			}
			if threeSum > target {
				r--
			} else {
				l++
			}
		}
	}
	return res
}
