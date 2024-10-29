package nsum

import (
	"sort"
)

func threeSumMulti(arr []int, target int) int {
	length := len(arr)
	sort.Ints(arr)
	count := 0

	const MOD = 1000000000 + 7
	for i := 0; i < length-2; i++ {
		l, r := i+1, length-1
		for l < r {
			threeSum := arr[i] + arr[l] + arr[r]
			if threeSum < target {
				l++
			} else if threeSum > target {
				r--
			} else {
				if arr[r] == arr[l] {
					count += (r - l) * (r - l + 1) / 2
					break
				} else {
					left, right := 1, 1
					for l+1 < r && arr[l] == arr[l+1] {
						left++
						l++
					}
					for r-1 > l && arr[r] == arr[r-1] {
						right++
						r--
					}
					count += left * right
					l++
					r--
				}
			}
		}
	}

	return count % MOD
}
