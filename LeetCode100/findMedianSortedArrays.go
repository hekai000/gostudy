package leetcode100

import "math"

func getKth(nums1 []int, start1 int, end1 int, nums2 []int, start2 int, end2 int, k int) float64 {
	len1, len2 := end1-start1+1, end2-start2+1
	if len1 > len2 {
		return getKth(nums2, start2, end2, nums1, start1, end1, k)
	}
	if len1 == 0 {
		return float64(nums2[start2+k-1])
	}
	if k == 1 {
		return float64(math.Min(float64(nums1[start1]), float64(nums2[start2])))
	}
	i, j := start1+int(math.Min(float64(len1), math.Floor(float64(k)/2)))-1, start2+int(math.Min(float64(len2), math.Floor(float64(k)/2)))-1
	if nums1[i] > nums2[j] {
		return getKth(nums1, start1, end1, nums2, j+1, end2, k-(j-start2+1))
	} else {
		return getKth(nums1, i+1, end1, nums2, start2, end2, k-(i-start1+1))
	}
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n, m := len(nums1), len(nums2)
	if n == 0 && m == 0 {
		return 0
	}
	left, right := (n+m+1)/2, (n+m+2)/2
	return (getKth(nums1, 0, n-1, nums2, 0, m-1, left) + getKth(nums1, 0, n-1, nums2, 0, m-1, right)) * 0.5

}
