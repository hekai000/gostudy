package twopointers

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	memory := make(map[byte]int)
	left := 0
	res := 0
	for index, c := range s {
		memory[byte(c)]++
		if len(memory) > k {
			memory[s[left]]--
			if memory[s[left]] == 0 {
				delete(memory, s[left])
			}
			left++
		}
		res = max(res, index-left+1)
	}
	return len(s) - left
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
