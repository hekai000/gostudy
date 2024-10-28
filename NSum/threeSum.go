package nsum

import "sort"

//Solution 1, two pointers
// func threeSum(nums []int) [][]int {
// 	length := len(nums)
// 	threeSumResult := [][]int{}
// 	if length < 3 {
// 		return [][]int{}
// 	}
// 	sort.Ints(nums)
// 	for i := 0; i < length-2; i++ {
// 		if i > 0 && nums[i] == nums[i-1] {
// 			continue
// 		}
// 		l, r := i+1, length-1
// 		for l < r {
// 			sum := nums[i] + nums[l] + nums[r]
// 			if sum == 0 {
// 				result := []int{nums[i], nums[l], nums[r]}
// 				// fmt.Println(result)
// 				threeSumResult = append(threeSumResult, result)
// 				l++
// 				r--
// 				for l < r && nums[l] == nums[l-1] {
// 					l++
// 				}
// 				for l < r && nums[r] == nums[r+1] {
// 					r--
// 				}
// 			} else if sum < 0 {
// 				l++
// 			} else {
// 				r--
// 			}
// 		}
// 	}
// 	return threeSumResult
// }

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		myTwoSum(nums, l, r, -nums[i], &res)
	}
	return res
}

func myTwoSum(nums []int, left int, right int, target int, res *[][]int) {
	memory := make(map[int]int)
	for i := left; i <= right; i++ {
		if _, ok := memory[target-nums[i]]; ok {
			result := []int{-target, target - nums[i], nums[i]}
			*res = append(*res, result)
			for i < right && nums[i] == nums[i+1] {
				i++
			}
		}
		memory[nums[i]] = 1
	}
}
