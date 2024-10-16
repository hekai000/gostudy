package leetcode100

func palindRome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
func longestPalindrome(s string) string {
	length := len(s)
	longestP := ""
	for i := 0; i < length; i++ {
		p1 := palindRome(s, i, i)
		p2 := palindRome(s, i, i+1)
		if len(p1) > len(longestP) {
			longestP = p1
		}
		if len(p2) > len(longestP) {
			longestP = p2
		}
	}
	return longestP
}
