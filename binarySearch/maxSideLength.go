package binarysearch

/*
预备知识
本题需要用到一些二维前缀和（Prefix Sum）的知识，它是一维前缀和的延伸：

设二维数组 A 的大小为 m * n，行下标的范围为 [1, m]，列下标的范围为 [1, n]。

数组 P 是 A 的前缀和数组，等价于 P 中的每个元素 P[i][j]：

如果 i 和 j 均大于 0，那么 P[i][j] 表示 A 中以 (1, 1) 为左上角，(i, j) 为右下角的矩形区域的元素之和；

如果 i 和 j 中至少有一个等于 0，那么 P[i][j] 也等于 0。

数组 P 可以帮助我们在 O(1) 的时间内求出任意一个矩形区域的元素之和。
具体地，设我们需要求和的矩形区域的左上角为 (x1, y1)，右下角为 (x2, y2)，
则该矩形区域的元素之和可以表示为：
sum = A[x1..x2][y1..y2]

	= P[x2][y2] - P[x1 - 1][y2] - P[x2][y1 - 1] + P[x1 - 1][y1 - 1]

P[i][j] = P[i - 1][j] + P[i][j - 1] - P[i - 1][j - 1] + A[i][j]
*/
func maxSideLength(mat [][]int, threshold int) int {

	var right int
	var sum int
	m, n := len(mat), len(mat[0])
	P := make([][]int, m+1)
	for i := range P {
		P[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			P[i][j] = P[i-1][j] + P[i][j-1] - P[i-1][j-1] + mat[i-1][j-1]
		}
	}

	if m < n {
		right = m
	} else {
		right = n
	}
	left := 0
	ans := 0
	for left <= right {
		mid := left + (right-left)/2
		find := false
		for i := 1; i <= m-mid+1; i++ {
			for j := 1; j <= n-mid+1; j++ {
				sum = getSum(P, i, j, i+mid-1, j+mid-1)
				if sum <= threshold {
					find = true
				}
			}
		}
		if find {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

func getSum(P [][]int, x1 int, y1 int, x2 int, y2 int) int {

	return P[x2][y2] - P[x1-1][y2] - P[x2][y1-1] + P[x1-1][y1-1]
}
