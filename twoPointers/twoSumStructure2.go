package twopointers

type TwoSum2 struct {
	sums map[int]int
	nums []int
}

func TwoSum2Constructor() TwoSum2 {
	return TwoSum2{map[int]int{}, []int{}}
}

func (this *TwoSum2) AddTwoSum2(number int) {
	this.nums = append(this.nums, number)

	for i := 0; i < len(this.nums)-1; i++ {
		sum := this.nums[i] + this.nums[len(this.nums)-1]
		this.sums[sum] = 1
	}

}

func (this *TwoSum2) FindTwoSum2(target int) bool {
	return this.sums[target] == 1
}
