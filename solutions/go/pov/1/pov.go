package pov

import "slices"

type Tree struct {
	// Add the needed fields here
	value    string
	children []*Tree
	parent   *Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	node := &Tree{
		value:    value,
		children: make([]*Tree, 0, len(children)),
		parent:   nil,
	}
	for _, child := range children {
		child.parent = node // 反向绑定：子节点的父指向当前节点
		node.children = append(node.children, child)
	}
	return node
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	return tr.value
}

// Children returns a slice containing the children of a tree.
// There is no need to sort the elements in the result slice,
// they can be in any order.
func (tr *Tree) Children() []*Tree {
	return tr.children
}

// String describes a tree in a compact S-expression format.
// This helps to make test outputs more readable.
// Feel free to adapt this method as you see fit.
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

// POV problem-specific functions

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	// 【改动1】先在原树中定位目标节点，找不到直接返回 nil（原来错误地返回 tr，吞掉了错误）
	node := tr.findByValue(from)
	if node == nil {
		return nil
	}
	// 【改动2】复制整棵树再翻转，原树保持不变
	// （原来的 reverseTree 直接改 children/parent，调用一次就把原树破坏了；
	//   因为要沿 parent 往上翻转，所以必须整棵复制）
	dup := copyTree(tr)
	return reverseTree(dup.findByValue(from), nil)
}

// PathTo returns the shortest path between two nodes in the tree.
// 【改动3】PathTo 重写：先校验两端节点是否存在，再从 from 向上收集整条祖先链，
// 从 to 向上找到第一个落在祖先链上的节点（即最近公共祖先 LCA），
// 拼成 from -> ... -> LCA -> ... -> to 的最短路径。
// 原实现的问题：node1 = node1.parent 未判空会 panic；
// to 不在 from 子树内时 path2 每轮被清空导致结果错误/死循环；
// path2 = append(path1, ...) 复用底层数组造成结果互相污染。
func (tr *Tree) PathTo(from, to string) []string {
	node1 := tr.findByValue(from)
	node2 := tr.findByValue(to)
	// 【改动3-a】任一端不存在（原来会 panic 或返回错误结果）直接返回 nil
	if node1 == nil || node2 == nil {
		return nil
	}

	// 【改动3-b】从 from 向上收集整条祖先链，并记录每个祖先在链上的下标
	up := []string{}
	ancestors := map[*Tree]int{}
	for n := node1; n != nil; n = n.parent {
		ancestors[n] = len(up)
		up = append(up, n.value)
	}

	// to 是 from 的祖先：直接取 up 的前缀，已经是正确顺序
	if i, ok := ancestors[node2]; ok {
		return slices.Clone(up[:i+1])
	}

	// 【改动3-c】从 to 向上回溯，命中祖先链的节点即 LCA；
	// down 收集到的是 to -> ... -> LCA 的孩子，反转后拼在 LCA 之后
	down := []string{}
	lcaIdx := -1
	for n := node2; n != nil; n = n.parent {
		if i, ok := ancestors[n]; ok {
			lcaIdx = i
			break
		}
		down = append(down, n.value)
	}
	if lcaIdx < 0 {
		return nil
	}
	slices.Reverse(down)
	return append(up[:lcaIdx+1:lcaIdx+1], down...)
}

func (tr *Tree) findByValue(value string) *Tree {
	if tr.value == value {
		return tr
	}
	for _, cur := range tr.children {
		if res := cur.findByValue(value); res != nil {
			return res
		}
	}
	return nil
}

// 【改动4】reverseTree 改为包级函数并改成迭代实现。
// 原递归实现有两个问题：
//
//	a) 原来挂在 *Tree 上（`node.reverseTree(...)`），容易误读成在 node 上操作；
//	b) 递归返回的是「最顶层节点」（也就是原树的根），而不是翻转后的新根。
//	   例如 parent(x, sibling) 从 x 翻转，递归结束时返回的是 parent，导致 FromPov
//	   返回了原树根。现在用 newRoot 显式保存起点。
func reverseTree(node *Tree, parent *Tree) *Tree {
	newRoot := node // 翻转后的新根就是传入的起点节点
	for n := node; n != nil; {
		curParent := n.parent
		// 删掉朝 parent 方向的那个孩子（改由 parent 持有 n），避免出现环
		n.children = slices.DeleteFunc(n.children, func(cur *Tree) bool {
			return cur == parent
		})
		n.parent = parent
		if curParent == nil {
			break // 已经翻到原树的根，结束
		}
		// 原父节点降级为 n 的孩子，继续沿 parent 往上翻转
		n.children = append(n.children, curParent)
		parent, n = n, curParent
	}
	return newRoot
}

// 【改动5】新增 copyTree 深拷贝，用于让 FromPov 不破坏原树。
// 不能只复制目标节点的子树，因为翻转需要沿 parent 指针一路走到根，
// 那些祖先节点也必须一并复制。
func copyTree(tr *Tree) *Tree {
	if tr == nil {
		return nil
	}
	node := &Tree{value: tr.value}
	for _, child := range tr.children {
		dupChild := copyTree(child)
		dupChild.parent = node // 反向绑定副本的父子关系
		node.children = append(node.children, dupChild)
	}
	return node
}
