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
	for j := range uf.rank {
		uf.rank[j] = 1
	}
}

func (uf *UnionFind) Find(p int) int {
	// root := p
	// //查找根
	// for root != uf.parent[root] {
	// 	root = uf.parent[root]
	// }
	// //完全压缩
	// for p != uf.parent[p] {
	// 	temp := uf.parent[p]
	// 	uf.parent[p] = root
	// 	p = temp
	// }
	//完全压缩
	if uf.parent[p] != p {
		uf.parent[p] = uf.Find(uf.parent[p])
	}
	return uf.parent[p]
	// return root
}

func (uf *UnionFind) Union(p, q int) {
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)
	if pRoot == qRoot {
		return
	}
	if uf.rank[qRoot] > uf.rank[pRoot] {
		uf.parent[pRoot] = qRoot
		uf.rank[qRoot] += uf.rank[pRoot]
	} else if uf.rank[pRoot] > uf.rank[qRoot] {
		uf.parent[qRoot] = pRoot
	} else {
		uf.parent[qRoot] = pRoot
		uf.rank[pRoot] += 1
	}
	uf.count--
}

func (uf *UnionFind) TotalCount() int {
	return uf.count
}
