package nsum

import "sort"

// 使用双指针解法解twoSum问题
func twoSum(nums []int, target int) []int {
	if len(nums) <= 1 {
		return []int{}
	}
	nums2 := make([]int, len(nums))
	copy(nums2, nums)

	sort.Ints(nums2)
	left, right := 0, len(nums2)-1
	res := make([]int, 2)

	for left < right && nums2[left]+nums2[right] != target {
		if nums2[left]+nums2[right] > target {
			right--
		} else {
			left++
		}
	}
	for index1, num := range nums {
		if num == nums2[left] {
			res[0] = index1
		}
	}
	for index2, num := range nums {
		if num == nums2[right] && index2 != res[0] {
			res[1] = index2
			break
		}
	}
	return res
}
