package leetcode100

func splitArray(nums []int, k int) int {
	sum, maxv := 0, 0
	var mid int
	for _, num := range nums {
		if num > maxv {
			maxv = num
		}
		sum += num
	}
	left, right := maxv, sum

	for left < right {
		mid = left + (right-left)/2
		if test(nums, mid, k) {
			left = mid + 1
		} else {
			right = mid
		}

	}
	return left
}

func test(nums []int, mid, k int) bool {
	count, asum := 1, 0
	for _, num := range nums {
		if asum+num > mid {
			asum = num
			count++
		} else {
			asum += num
		}
	}
	return count > k
}
