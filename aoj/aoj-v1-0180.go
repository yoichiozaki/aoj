// Find the minimum spanning tree by the Kruskal algorithm
// -------------------------------------------------------
// Kruskal(G):
//     A = ∅
//     foreach v ∈ G.V:
//        MAKE-SET(v)
//     foreach (u, v) in G.E ordered by weight(u, v), increasing:
//         if FIND-SET(u) ≠ FIND-SET(v):
//             A = A ∪ {(u, v)}
//             UNION(u, v)
//     return A
//
// Ref: https://en.wikipedia.org/wiki/Kruskal%27s_algorithm

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)
var (
	matrix [][]int
)

type Edge struct {
	a, b, cost, added int
}
type ByCost []Edge

var edges ByCost

func (a ByCost) Len() int           { return len(a) }
func (a ByCost) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCost) Less(i, j int) bool { return a[i].cost < a[j].cost }

func main() {
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(buf[0])
		m, _ := strconv.Atoi(buf[1])
		if n == 0 && m == 0 {
			break
		}
		matrix = make([][]int, n)
		for i := range matrix {
			matrix[i] = make([]int, n)
		}
		for i := 0; i < m; i++ {
			scanner.Scan()
			buf_ := strings.Split(scanner.Text(), " ")
			a, _ := strconv.Atoi(buf_[0])
			b, _ := strconv.Atoi(buf_[1])
			cost, _ := strconv.Atoi(buf_[2])
			matrix[a][b] = cost
			matrix[b][a] = cost
		}
		sum := Kruskal(n)
		fmt.Println(sum)
		if len(edges) > 0 {
			edges = []Edge{}
		}
	}
}

func Kruskal_func(size, edge_num int) int {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if matrix[i][j] > 0 {
				tmp := Edge{i, j, matrix[i][j], 0}
				edges = append(edges, tmp)
			}
		}
	}
	sort.Sort(ByCost(edges))
	UF_Init(edge_num)
	for i := 0; i < edge_num; i++ {
		a := edges[i].a
		b := edges[i].b
		if UF_Find(a) != UF_Find(b) {
			edges[i].added = 1
			UF_Union(UF_Find(a), UF_Find(b))
		}
	}
	sum := 0
	for i := 0; i < edge_num; i++ {
		if edges[i].added != 0 {
			// log.Printf("edge(%d,%d) is added\n", edges[i].a, edges[i].b)
			sum += edges[i].cost
		}
	}
	return sum
}

func Kruskal(size int) int {
	edge_num := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if matrix[i][j] > 0 {
				edge_num++
			}
		}
	}
	return Kruskal_func(size, edge_num)
}

/*** Union-Find functions ***/
var UF_Parent []int
var UF_Rank []int

func UF_Init(n int) {
	UF_Parent = make([]int, n)
	UF_Rank = make([]int, n)
	for i := 0; i < n; i++ {
		UF_Parent[i] = i
		UF_Rank[i] = 0
	}
}
func UF_Union(a, b int) {
	if UF_Rank[a] < UF_Rank[b] {
		UF_Parent[a] = b
	} else {
		UF_Parent[b] = a
	}
	if UF_Rank[a] == UF_Rank[b] {
		UF_Rank[a]++
	}
}
func UF_Find(a int) int {
	if UF_Parent[a] == a {
		return a
	}
	return UF_Find(UF_Parent[a])
}

/*** Union-Find functions ***/
