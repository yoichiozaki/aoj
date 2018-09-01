package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

type UnionFind struct {
	parents []int
}

func (u *UnionFind) init(N int) {
	u.parents = make([]int, N)
	for i := range u.parents {
		u.parents[i] = i // ノードiの親はノードi自身であるとして初期化
	}
}
func (u *UnionFind) root(x int) int { // ノードxの親を再帰で手繰ることで得る
	if u.parents[x] == x {
		return x
	}
	u.parents[x] = u.root(u.parents[x])
	return u.parents[x]
}
func (u *UnionFind) unite(x, y int) { // ノードxとノードyの属する木が異なればそれらを統合する
	rx := u.root(x)
	ry := u.root(y)
	if rx == ry {
		return
	}
	u.parents[rx] = ry
}
func (u *UnionFind) same(x, y int) bool { // ノードxとノードyが同じ木に属する華道家の判定
	rx := u.root(x)
	ry := u.root(y)
	return rx == ry
}

func main() {
	scanner.Scan()
	buf := strings.Split(scanner.Text(), " ")
	N, _ := strconv.Atoi(buf[0])
	Q, _ := strconv.Atoi(buf[1])
	var tree UnionFind
	tree.init(N)
	for i := 0; i < Q; i++ {
		scanner.Scan()
		buf_ := strings.Split(scanner.Text(), " ")
		P, _ := strconv.Atoi(buf_[0])
		A, _ := strconv.Atoi(buf_[1])
		B, _ := strconv.Atoi(buf_[2])
		switch P {
		case 0: // Unite
			tree.unite(A, B)
		case 1: // Check
			if tree.same(A, B) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}

}
