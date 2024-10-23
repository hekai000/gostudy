package binarysearch

import "sort"

func maxDistance(position []int, m int) int {
	if m == 0 {
		return 0
	}
	slice := position[:]
	sort.Ints(slice)
	left, right, ans := 1, position[len(position)-1]-position[0], 0
	for left <= right {
		mid := (left + right) / 2
		if check(slice, mid, m) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
func check(position []int, mid int, m int) bool {
	count := 1
	last := position[0]
	for i := 1; i < len(position); i++ {
		if position[i]-last >= mid {
			count++
			last = position[i]
			if count >= m {
				return true
			}
		}
	}
	return false
}
