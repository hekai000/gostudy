package twopointers

func maxArea(height []int) int {
	left, right, res := 0, len(height)-1, 0

	for left < right {
		width := right - left
		minHigh := min(height[left], height[right])
		res = max(res, width*minHigh)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
