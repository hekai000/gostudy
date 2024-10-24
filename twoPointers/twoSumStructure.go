package twopointers

type TwoSum struct {
	cnt map[int]int
}

func Constructor() TwoSum {
	return TwoSum{map[int]int{}}
}

func (this *TwoSum) Add(number int) {
	this.cnt[number] += 1
}

func (this *TwoSum) Find(value int) bool {
	for x, v := range this.cnt {
		y := value - v
		if _, ok := this.cnt[y]; ok && (x != y || v > 1) {
			return true
		}
	}
	return false
}
