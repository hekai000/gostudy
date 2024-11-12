package nsum

import "sort"

func my2Sum(nums []int, target int, start int) [][]int {
	ans := [][]int{}
	lo, hi := start, len(nums)-1
	for lo < hi {
		sum := nums[lo] + nums[hi]
		if sum == target {
			ans = append(ans, []int{nums[lo], nums[hi]})
			lo++
			hi--
			for lo < hi && nums[lo] == nums[lo-1] {
				lo++
			}
			for lo < hi && nums[hi] == nums[hi+1] {
				hi--
			}
		} else if sum < target {
			lo++
		} else {
			hi--

		}

	}
	return ans
}
func kSum(nums []int, k int, target int, start int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	avg := target / k
	if start == len(nums) {
		return ans
	}
	if nums[start] > avg || nums[len(nums)-1] < avg {
		return ans
	}
	if k == 2 {
		return my2Sum(nums, target, start)
	}
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > avg {
			break
		}
		for _, v := range kSum(nums, k-1, target-nums[i], i+1) {
			ans = append(ans, append([]int{nums[i]}, v...))

		}
	}
	return ans

}
