package util

import (
	"fmt"
	"math/bits"
)

type UnionFind []int

func (u UnionFind) Find(x int) int {
	if u[x] != x {
		u[x] = u.Find(u[x])
	}
	return u[x]
}

func (u UnionFind) Union(x, y int) {
	u[u.Find(y)] = u.Find(x)
}

func (u UnionFind) NGroup() int {
	res := 0
	for i, v := range u {
		if i == v {
			res++
		}
	}
	return res
}

// type UnionFind[T comparable] map[T]T

// func (u UnionFind[T]) Find(x T) T {
// 	if v, ok := u[x]; !ok {
// 		u[x] = x
// 	} else if v != x {
// 		u[x] = u.Find(v)
// 	}
// 	return u[x]
// }

// func (u UnionFind[T]) Union(x, y T) {
// 	u[u.Find(y)] = u.Find(x)
// }

type Stack[T any] struct {
	data []T
	i    int
}

func (s *Stack[T]) Top() T {
	return s.data[s.i-1]
}

func (s *Stack[T]) Len() int {
	return s.i
}

func (s *Stack[T]) Push(v T) {
	if s.i == len(s.data) {
		s.data = append(s.data, v)
	} else {
		s.data[s.i] = v
	}
	s.i += 1
}

func (s *Stack[T]) Pop() T {
	s.i -= 1
	if s.i < 0 {
		panic("Invalid stack pop: stack is empty")
	}
	return s.data[s.i]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: []T{}, i: 0}
}

// type Deque[T any] struct {
// 	data []T
// 	i, j int
// }

type Counter[T comparable] struct {
	M map[T]int
}

func (c *Counter[T]) Init() {
	if c.M == nil {
		c.M = map[T]int{}
	}
}

func (c *Counter[T]) Update(l []T) {
	c.Init()
	for _, v := range l {
		c.M[v] += 1
	}
}

type Pair [2]int

// PairList represents a slice of pairs.
type PairList []Pair

// Len returns the length of the pair list.
func (p PairList) Len() int { return len(p) }

// Less compares two pairs by their keys.
func (p PairList) Less(i, j int) bool {
	return p[i][0] < p[j][0] || (p[i][0] == p[j][0] && p[i][1] < p[j][1])
}

// Swap swaps two pairs in the pair list.
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SegNode struct {
	Lo, Hi      int
	Ac, V       int64
	Left, Right *SegNode
}

func CreateNode(acl, vl []int64) *SegNode {
	var create func(lo, hi int) *SegNode
	create = func(lo, hi int) *SegNode {
		if lo == hi {
			return &SegNode{lo, hi, acl[lo], vl[lo], nil, nil}
		}
		mid := (lo + hi) / 2
		res := &SegNode{lo, hi, 0, 0, create(lo, mid), create(mid+1, hi)}
		res.Fixv()
		return res
	}
	return create(0, len(acl)-1)
}

func (node *SegNode) Fixv() {
	dif := node.Right.Ac - node.Left.V
	if dif > 1 {
		node.Ac, node.V = inf, node.Right.V
	} else if dif == 1 {
		node.Ac, node.V = node.Left.Ac, node.Right.V
	} else {
		node.Ac, node.V = inf, node.Right.V+1
	}
}

func (node *SegNode) Update(i int, ac, v int64) {
	if node.Lo == i && node.Hi == i {
		node.Ac, node.V = ac, v
		return
	}
	if i <= node.Left.Hi {
		node.Left.Update(i, ac, v)
	} else {
		node.Right.Update(i, ac, v)
	}
	node.Fixv()
}

type Deque[T any] struct {
	Data       []T
	Start, End int
}

func (dq *Deque[T]) Len() int {
	return dq.End - dq.Start
}

func NewDeque[T any](l []T, capacity int) *Deque[T] {
	res := &Deque[T]{}
	if capacity < 2 {
		capacity = 2
	}
	res.Start = 0
	res.End = len(l)
	if len(l) >= capacity {
		res.Data = l
	} else {
		res.Data = make([]T, capacity)
		copy(res.Data, l)
	}
	// fmt.Println(res)
	return res
}

func (dq *Deque[T]) upgradeFull() {
	if dq.Len() < len(dq.Data) {
		return
	}
	// fmt.Println("upgrade", dq)
	data := append(dq.Data[dq.Start:], dq.Data[:dq.Start]...)
	dq.Start = 0
	dq.End = len(data)
	dq.Data = make([]T, 2*len(data)+1)
	copy(dq.Data, data)
}

func (dq *Deque[T]) First() T {
	// fmt.Println(dq)
	return dq.Data[dq.Start]
}
func (dq *Deque[T]) Last() T {
	// fmt.Println(dq)
	j := dq.End - 1
	if j >= len(dq.Data) {
		j -= len(dq.Data)
	}
	return dq.Data[j]
}

func (dq *Deque[T]) At(i int) T {
	i += dq.Start
	if i >= len(dq.Data) {
		i -= len(dq.Data)
	}
	return dq.Data[i]
}

func (dq *Deque[T]) AppendLeft(v T) {
	dq.upgradeFull()
	if dq.Start == 0 {
		dq.Start, dq.End = dq.Start+len(dq.Data), dq.End+len(dq.Data)
	}
	dq.Start -= 1
	dq.Data[dq.Start] = v
}

func (dq *Deque[T]) Append(v T) {
	dq.upgradeFull()
	j := dq.End
	if j >= len(dq.Data) {
		j -= len(dq.Data)
	}
	dq.Data[j] = v
	dq.End += 1
}

func (dq *Deque[T]) PopLeft() T {
	res := dq.First()
	dq.Start += 1
	if dq.Start == len(dq.Data) {
		dq.Start, dq.End = dq.Start-len(dq.Data), dq.End-len(dq.Data)
	}
	return res
}

func (dq *Deque[T]) Pop() T {
	res := dq.Last()
	dq.End -= 1
	return res
}

func (dq Deque[T]) String() string {
	res := make([]T, dq.Len())
	for i := 0; i < dq.Len(); i++ {
		res[i] = dq.At(i)
	}
	return fmt.Sprint(res)
}

type QueryRange interface {
	Query(i, j int) (min_index int)
}

type rmqQuery struct {
	query func(int, int) int
}

func (q *rmqQuery) Query(i int, j int) (min_index int) {
	return q.query(i, j)
}

func RangeMinimumQuery(data []int) QueryRange {
	n := len(data)
	nbit := bits.Len(uint(n))
	res := make([][]int, nbit)
	res[0] = makeRange(0, n)
	choose_min := func(i, j int) int {
		if data[i] <= data[j] {
			return i
		} else {
			return j
		}
	}
	for j := 1; j < nbit; j++ {
		added := 1 << (j - 1)
		prev, cur := res[j-1], make([]int, n+1-2*added)
		for i := 0; i < len(cur); i++ {
			cur[i] = choose_min(prev[i], prev[i+added])
		}
		res[j] = cur
	}
	query := func(i, j int) (min_index int) {
		nbit := bits.Len(uint(j-i+1)) - 1
		return choose_min(res[nbit][i], res[nbit][j+1-(1<<nbit)])
	}

	return &rmqQuery{query}
	// return the query range
}

type FenwickTree struct {
	A    []int
	Tree []int
}

func (t *FenwickTree) Update(i, v int) {
	if t.A[i] == v {
		return
	}
	i += 1
	added := v - t.A[i]
	t.A[i] += added
	for i < len(t.Tree) {
		t.Tree[i] += added
		i += i & -i
	}
}

func (t *FenwickTree) Sum(i int) int {
	res := 0
	for i > 0 {
		res += t.Tree[i]
		i -= i & -i
	}
	return res
}

func (t *FenwickTree) Range(i, j int) int {
	return t.Sum(j+1) - t.Sum(i)
}

func NewFenwickTree(a []int) *FenwickTree {
	res := &FenwickTree{}
	res.A = a
	res.Tree = append([]int{0}, a...)
	for i := 1; i < len(res.Tree); i++ {
		j := i + (i & -i)
		if j < len(res.Tree) {
			res.Tree[j] += res.Tree[i]
		}
	}
	return res
}
