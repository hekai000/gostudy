package segmentCountTree

type SegmentCountTree struct {
	data, tree  []int
	left, right int
	merge       func(i, j int) int
}

func (st *SegmentCountTree) Init(nums []int, oper func(i, j int) int) {
	st.merge = oper
	data, tree := make([]int, len(nums)), make([]int, 4*len(nums))
	for i := 0; i < len(nums); i++ {
		data[i] = nums[i]
	}
	st.data, st.tree = data, tree

}
func (st *SegmentCountTree) leftChild(index int) int {
	return 2*index + 1
}
func (st *SegmentCountTree) rightChild(index int) int {
	return 2*index + 2
}

func (st *SegmentCountTree) buildSegmentCountTree(treeIndex, left, right int) {
	if left == right {
		st.tree[treeIndex] = st.data[left]
		return
	}
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
	st.buildSegmentCountTree(leftTreeIndex, left, midTreeIndex)
	st.buildSegmentCountTree(rightTreeIndex, midTreeIndex+1, right)
	st.tree[treeIndex] = st.merge(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

func (st *SegmentCountTree) Query(left, right int) int {
	if len(st.data) > 0 {
		return st.queryInCountTree(0, 0, len(st.data)-1, left, right)
	}
	return 0
}

func (st *SegmentCountTree) queryInCountTree(treeIndex, left, right, queryLeft, queryRight int) int {
	if queryLeft > st.data[right] || queryRight < st.data[left] {
		return 0
	}
	if queryLeft <= st.data[left] && st.data[right] <= queryRight || left == right {
		return st.tree[treeIndex]
	}
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
	return st.queryInCountTree(leftTreeIndex, left, midTreeIndex, queryLeft, queryRight) + st.queryInCountTree(rightTreeIndex, midTreeIndex+1, right, queryLeft, queryRight)
}

func (st *SegmentCountTree) Update(value int) {
	if len(st.data) > 0 {
		st.updateInCountTree(0, 0, len(st.data)-1, value)
	}
}

func (st *SegmentCountTree) updateInCountTree(treeIndex, left, right, value int) {
	if value >= st.data[left] && value <= st.data[right] {
		st.tree[treeIndex]++
		if left == right {
			return
		}
	}
	midTreeIndex, leftTreeIndex, rightTreeIndex := left+(right-left)>>1, st.leftChild(treeIndex), st.rightChild(treeIndex)
	st.updateInCountTree(leftTreeIndex, left, midTreeIndex, value)
	st.updateInCountTree(rightTreeIndex, midTreeIndex+1, right, value)
}
