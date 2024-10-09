package util

import (
	"fmt"
)

type Node[T any] struct {
	V           T
	C           int
	Left, Right *Node[T]
}

func (n *Node[T]) Fix() {
	n.C = 1
	if n.Left != nil {
		n.C += n.Left.C
	}
	if n.Right != nil {
		n.C += n.Right.C
	}
}

func (n *Node[T]) Index() int {
	if n.Left != nil {
		return n.Left.C
	}
	return 0
}

func (n *Node[T]) Swap(x *Node[T]) {
	n.V, n.Left, n.Right, x.V, x.Left, x.Right = x.V, x.Left, x.Right, n.V, n.Left, n.Right
}

func (n *Node[T]) LeftRotate() *Node[T] {
	if n.Right == nil {
		return n
	}
	x := n.Right
	n.Swap(x)
	x.Right, n.Left = n.Left, x
	x.Fix()
	n.Fix()
	return n

}

func (n *Node[T]) RightRotate() *Node[T] {
	if n.Left == nil {
		return n
	}
	x := n.Left
	n.Swap(x)
	x.Left, n.Right = n.Right, x
	x.Fix()
	n.Fix()
	return n
}

func (n *Node[T]) Rebalance() *Node[T] {
	if n.C <= 2 {
		return n
	}
	getc := func(n *Node[T]) int {
		if n == nil {
			return 0
		}
		return n.C
	}
	if getc(n.Left)*3 < n.C-1 {
		if getc(n.Left) >= getc(n.Right.Right) {
			n.Right.RightRotate()
		}
		n.LeftRotate()
	} else if getc(n.Right)*3 < n.C-1 {
		if getc(n.Right) >= getc(n.Left.Left) {
			n.Left.LeftRotate()
		}
		n.RightRotate()
	}
	return n
}

func (n *Node[T]) ToMap() map[string]any {
	res := map[string]any{
		"_a": n.V,
		"_c": n.C,
	}
	if n.Left != nil {
		res["left"] = n.Left.ToMap()
	}
	if n.Right != nil {
		res["right"] = n.Right.ToMap()
	}
	return res
}

func (n *Node[T]) PrettyPrint() {
	// m := n.ToMap()
	// res, _ := json.MarshalIndent(m, "", " ")
	// fmt.Println(string(res))
	d := [][]string{}
	var dp func(n *Node[T], i int)
	dp = func(n *Node[T], i int) {
		if n == nil {
			return
		}
		if i == len(d) {
			d = append(d, []string{})
		}
		d[i] = append(d[i], fmt.Sprintf("(%v,%v)", n.V, n.C))
		dp(n.Left, i+1)
		dp(n.Right, i+1)
	}
	dp(n, 0)
	for _, r := range d {
		fmt.Println(r)
	}

}

type SplayTree[T any] struct {
	LessEq func(a, b T) bool
	Root   *Node[T]
}

func NewSplayTree[T any](lessEq func(a, b T) bool) *SplayTree[T] {
	return &SplayTree[T]{
		LessEq: lessEq, Root: nil,
	}
}

func (tree *SplayTree[T]) Len() int {
	if tree.Root == nil {
		return 0
	}
	return tree.Root.C
}

func (tree *SplayTree[T]) splay(path []*Node[T]) {
	for len(path) > 2 {
		grandpa, p, x := path[len(path)-3], path[len(path)-2], path[len(path)-1]
		if grandpa.Left == p {
			if p.Left == x {
				grandpa.RightRotate()
				grandpa.RightRotate()
			} else {
				p.LeftRotate()
				grandpa.RightRotate()
			}
		} else {
			if p.Right == x {
				grandpa.LeftRotate()
				grandpa.LeftRotate()
			} else {
				p.RightRotate()
				grandpa.LeftRotate()
			}
		}
		path = path[:len(path)-2]
	}
	if len(path) == 0 {
		return
	} else if len(path) == 1 {
		tree.Root = path[0]
	} else if len(path) == 2 {
		tree.Root = path[0]
		if path[0].Left == path[1] {
			tree.Root.RightRotate()
		} else {
			tree.Root.LeftRotate()
		}
	}

	// if tree.Root.Left != nil {
	// 	tree.Root.Left.Rebalance()
	// }
	// if tree.Root.Right != nil {
	// 	tree.Root.Right.Rebalance()
	// }

}

func (tree *SplayTree[T]) Add(v T) {
	node, path := tree.Root, []*Node[T]{}
	for node != nil {
		path = append(path, node)
		if tree.LessEq(v, node.V) {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	node = &Node[T]{v, 1, nil, nil}
	if len(path) > 0 {
		last := path[len(path)-1]
		if tree.LessEq(v, last.V) {
			last.Left = node
		} else {
			last.Right = node
		}
		last.Fix()
	}
	path = append(path, node)
	tree.splay(path)
}

func (tree *SplayTree[T]) At(i int) T {
	node, path := tree.Root, []*Node[T]{}
	if node == nil || i < 0 || i >= node.C {
		panic(fmt.Errorf("invalid index %v", i))
	}
	for node != nil {
		path = append(path, node)
		if node.Index() == i {
			break
		} else if i < node.Index() {
			node = node.Left
		} else {
			i -= node.Index() + 1
			node = node.Right
		}
	}
	tree.splay(path)

	return tree.Root.V
}

func (tree *SplayTree[T]) Pop(i int) T {
	tree.At(i)
	res := tree.Root.V
	if tree.Root.Left == nil {
		tree.Root = tree.Root.Right
	} else if tree.Root.Right == nil {
		tree.Root = tree.Root.Left
	} else {
		right := tree.Root.Right
		tree.Root = tree.Root.Left
		tree.At(tree.Len() - 1)
		tree.Root.Right = right
		tree.Root.Fix()
	}
	return res
}

func (tree *SplayTree[T]) Delete(v T) (isdeleted bool) {
	node, path := tree.Root, []*Node[T]{}
	found := false
	for node != nil {
		path = append(path, node)
		if tree.LessEq(v, node.V) && tree.LessEq(node.V, v) {
			found = true
			break
		} else if tree.LessEq(v, node.V) {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	if !found {
		return false
	}
	tree.splay(path)
	tree.Pop(tree.Root.Index())
	return true
}

func (tree *SplayTree[T]) Bisect(v T) int {
	// Find index of value greater than or equal
	node, offset, cur := tree.Root, 0, 0
	if node == nil {
		return 0
	}
	cur = node.C
	for node != nil {
		if tree.LessEq(v, node.V) {
			cur = offset + node.Index()
			node = node.Left
		} else {
			offset += node.Index() + 1
			node = node.Right
		}
	}
	return cur
}
