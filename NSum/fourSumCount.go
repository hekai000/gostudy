package nsum

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	countAB := map[int]int{}
	ans := 0
	for _, v := range nums1 {
		for _, w := range nums2 {
			countAB[v+w]++
		}
	}

	for _, v := range nums3 {
		for _, w := range nums4 {
			ans += countAB[-v-w]
		}
	}
	return ans
}
