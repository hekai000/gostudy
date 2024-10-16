package leetcode100

func twoSum(nums []int, target int) []int {
	numsDict := make(map[int]int)
	result := []int{}
	for index, value := range nums {
		if i, ok := numsDict[target-value]; ok {
			return append(result, i, index)
		}
		numsDict[value] = index

	}
	return result
}
