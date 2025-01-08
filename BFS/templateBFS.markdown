# 二叉树的BFS

```golang BFS模板
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
func BFS(root *TreeNode) {
	if root == nil {
        return
    }

    q := []*TreeNode{}
    q = append(q, root)
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        fmt.Println(cur.Val)
        if cur.Left != nil {
            q = append(q, cur.Left)
        }
        if cur.Right != nil {
            q = append(q, cur.Right)
        }

    }
}
```

增加深度记录

```golang

func BFS(root *TreeNode) {
    if root == nil {
        return
    }
    q := []*TreeNode{root}
    depth := 1
    for len(q) > 0 {
        sz := len(q)
        for i = 0; i < sz; i++ {
            cur := q[0]
            q = q[1:]
            fmt.Printf("depth = %d, val = %d\n", depth, cur.Val)

        if cur.Left != nil {
            q = append(q, cur.Left)
        }
        if cur.Right != nil {
            q = append(q, cur.Rigth)
        }
        }
        depth++
    }
}
```

拓展
```golang BFS
    type State struct {
        node *TreeNode
        depth int
    }
    func BFS(root *TreeNode) {
        if root == nil {
            return
        }
        q := []*State{{root, 1}}
        for len(q) > 0 {
            cur := q[0]
            q = q[1:]
            // 访问 cur 节点，同时知道它的路径权重和
            fmt.Printf("depth = %d, val = %d\n", cur.depth, cur.node.Val)

            if cur.node.Left != nil {
                q = append(q, State{cur.node.Left, cur.depth + 1})
            }
            if cur.node.Right != nil {
                q = append(q, State{cur.node.Right, cur.depth + 1})
            }
        }
    }
```

补充
```golang BFS

    func BFS(root *TreeNode, target int) bool {
        if root == nil {
            return false
        }
        q := []*TreeNode{root}
        aset := make(map[*TreeNode]bool)
        aset[root] = true
        step := 0
        for len(q) > 0 {

            sz := len(q)
            for i:=0; i<sz; i++ {
                cur := q[0]
                q = q[1:]
                if cur.Val == target {
                    return true
                }
                if cur.Left != nil {
                    if aset[cur.Left] == false {
                        aset[cur.Left] = true
                        q = append(q, cur.Left)
                    }

                    
                }
                if cur.Right != nil {
                    if aset[cur.Right] == false {
                        aset[cur.Right] = true
                        q = append(q, cur.Right)
                    }
                    
                }
            }
            step++
        }
    }
```