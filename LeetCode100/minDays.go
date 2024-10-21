package leetcode100

import "math"

func minDays(bloomDay []int, m int, k int) int {
	if m > len(bloomDay)/k {
		return -1
	}
	if m == 0 {
		return 0
	}
	minDays, maxDays := math.MaxInt32, 0
	for _, ele := range bloomDay {
		if ele < minDays {
			minDays = ele
		}
		if ele > maxDays {
			maxDays = ele
		}
	}
	left, right, ans := minDays, maxDays, maxDays
	for left <= right {
		mid := left + (right-left)/2
		if checkBloomDay(bloomDay, m, k, mid) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	return ans
}

func checkBloomDay(bloomDay []int, m int, k int, days int) bool {
	count := 0
	countk := 0
	for i := 0; i < len(bloomDay); i++ {
		if bloomDay[i] <= days {
			countk++
			if countk == k {
				count++
				countk = 0
			}
		} else {
			countk = 0
		}
		if count == m {
			return true
		}
	}
	return false

}
