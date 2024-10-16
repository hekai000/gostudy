package leetcode100

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	slideWindow := make(map[byte]int, len(s))
	if length == 0 || length == 1 {
		return length
	}
	left, right, res := 0, 1, 0
	slideWindow[s[0]] = 0
	for right < length {
		if index, ok := slideWindow[s[right]]; !ok {
			slideWindow[s[right]] = right
			right++
		} else {

			for i := left; i <= index; i++ {
				delete(slideWindow, s[i])
			}
			left = index + 1

		}
		res = max(res, right-left)
	}
	return res
}
