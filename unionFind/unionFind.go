package unionfind

type UnionFind struct {
	parent []int
	rank   []int
	count  int
}

func (uf *UnionFind) Init(n int) {
	uf.count = n
	uf.parent = make([]int, n)
	uf.rank = make([]int, n)
	for i := range uf.parent {
		uf.parent[i] = i
	}
}

func (uf *UnionFind) Find(p int) int {
	root := p
	//查找根
	for root != uf.parent[root] {
		root = uf.parent[root]
	}
	//完全压缩
	for p != uf.parent[p] {
		temp := uf.parent[p]
		uf.parent[p] = root
		p = temp
	}
	return root
}
func (uf *UnionFind) Union(p, q int) {
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)
	if pRoot == qRoot {
		return
	}
	if uf.rank[qRoot] > uf.rank[pRoot] {
		uf.parent[pRoot] = qRoot
	} else {
		uf.parent[qRoot] = pRoot
		if uf.rank[pRoot] == uf.rank[qRoot] {
			uf.rank[pRoot]++
		}
	}
	uf.count--
}

func (uf *UnionFind) TotalCount() int {
	return uf.count
}
