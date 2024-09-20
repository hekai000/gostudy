package unionfind

type UnionFindCount struct {
	parent, count []int
	maxUnionCount int
}

func (uf *UnionFindCount) Init(n int) {
	uf.parent = make([]int, n)
	uf.count = make([]int, n)
	uf.maxUnionCount = 0
	for i := range uf.parent {
		uf.parent[i] = i
		uf.count[i] = 1
	}
}

func (uf *UnionFindCount) Find(p int) int {
	root := p
	for root != uf.parent[root] {
		root = uf.parent[root]
	}
	return root
}

func (uf *UnionFindCount) Union(p, q int) {
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)
	if pRoot == qRoot {
		return
	}
	if pRoot == len(uf.parent)-1 {
		pRoot, qRoot = qRoot, pRoot
	} else if qRoot == len(uf.parent)-1 {
		pRoot, qRoot = qRoot, pRoot
	} else if uf.count[qRoot] > uf.count[pRoot] {
		pRoot, qRoot = qRoot, pRoot
	}

	uf.maxUnionCount += max(uf.maxUnionCount, (uf.count[pRoot] + uf.count[qRoot]))
	uf.parent[qRoot] = pRoot
	uf.count[pRoot] += uf.count[qRoot]
}

func (uf *UnionFindCount) Count() []int {
	return uf.count
}

func (uf *UnionFindCount) MaxUnionCount() int {
	return uf.maxUnionCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
