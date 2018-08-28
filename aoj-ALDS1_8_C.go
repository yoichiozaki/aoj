
import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

type Node struct {
	value int
	left *Node
	right *Node
}


/****************************************************************
*関数insert():													*
*	 挿入される二分探索木のrootノードを表すポインタrootと				*
*	 挿入したいノードのvalueを引数に、適切な葉に新たなノードを挿入して、	*
*	 挿入された二分探索木のルートノードのポインタを返す関数				*
*****************************************************************/
func insert(root *Node, val int) *Node {

	// 二分探索木がまだ構築されていないときはここが呼び出される
	if root == nil {
		return &Node{val, nil,nil}
	}

	// すでにある二分探索木に対するノード（value: val）を挿入する
	// ここを抜けるとparentに挿入されるノードの親へのポインタが入っている
	// 変数nodeは、挿入すべき場所を探すための一時的なポインタ変数
	var parent *Node
	node := root
	for node != nil {
		parent = node
		if node.value > val {
			node = node.left
		} else {
			node = node.right
		}
	}

	// 新たなノードの挿入を実行
	if parent.value > val {
		parent.left = &Node{val, nil,nil}
	} else {
		parent.right = &Node{val, nil, nil}
	}

	// 二分探索木の根は変わらないのでそのまま返す
	return root
}

func printNodes(nodes []int) {
	// 木の走査によって得られたint型スライスを引数に、
	// 指定されたフォーマットに出力する関数
	writer := bufio.NewWriter(os.Stdout)
	for _, v := range nodes{
		writer.Write([]byte(fmt.Sprintf(" %d", v)))
	}
	writer.Write([]byte("\n"))
	writer.Flush()
}

func preorder(root *Node, memo []int) []int {
	// 引数rootが根となっている二分探索木に対して先行巡回によって
	// 各ノードをを舐めて、引数memoで与えられる空のスライスに書き込み、
	// 書き込まれた後のスライスを返す関数
	if root != nil {
		memo = append(memo, root.value)
		memo = preorder(root.left, memo)
		memo = preorder(root.right, memo)
	}
	return memo
}

func inorder(root *Node, memo []int) []int {
	// 引数rootが根となっている二分探索木に対して中間巡回によって
	// 各ノードをを舐めて、引数memoで与えられる空のスライスに書き込み、
	// 書き込まれた後のスライスを返す関数
	if root != nil {
		memo = inorder(root.left, memo)
		memo = append(memo, root.value)
		memo = inorder(root.right, memo)
	}
	return memo
}

/****************************************************************
*関数find():														*
*	 探索対象の二分探索木のrootノードを表すポインタrootと				*
*	 探したいノードのtargetを引数に、								*
*	 存在すればtrueを、なければfalseを返す関数						*
*****************************************************************/
func find(root *Node, target int) bool {
	switch {
	//　rootを根とする二分探索木が未構築
	case root == nil:
		return false
	case root.value == target:
		return true
	case root.value > target:
		return find(root.left, target)
	default:
		return find(root.right, target)
	}
}

/****************************************************************
*関数deleteNode():												*
*	 要素を削除される二分探索木のrootノードを表すポインタrootと			*
*	 削除対象のノードのvalueを引数に、ノードを削除する関数				*
* 																*
*	削除のパターン:												*
*		(1)削除対象のノードが子を0人持つ場合							*
*			そのまま対象ノードを削除する							*
*		(2)削除対象のノードが子を1人持つ場合							*
*			「削除対象ノードの親の子」を「削除対象のノードの子供」		*
*			にする												*
*		(3)削除対象のノードが子を2人持つ場合							*
*			削除対象ノードを根とするときの右部分木の最小値と削除対象	*
*			ノードを置き換える										*
*****************************************************************/
func deleteNode(root *Node, target int) {

	// targetをvalueとして持つノードを探す
	// このブロックを抜けた後のrootに格納されているポインタは
	// 削除対象のノードを指すポインタで、
	// parentは削除対象のノードであるrootの親ノードを指すポインタ
	parent := &Node{-1, nil, root}
	for root != nil {
		if root.value == target {
			break
		} else {
			parent = root
			if root.value < target {
				root = root.right
			} else {
				root = root.left
			}
		}
	}

	// 削除対象のノードが子を2人持つとき
	if root.left != nil && root.right != nil {
		min, minParent := findMinNode(root)
		root.value = min.value
		root, parent = min, minParent
	}

	// 削除対象のノードが子を0人持つとき
	if root.left == nil && root.right == nil {
		if parent.left == root {
			parent.left = nil
		} else {
			parent.right = nil
		}
	} else {
		// 削除対象のノードが子を1人以上持つとき
		next := root.right
		if next == nil {
			next = root.left
		}
		if parent.left == root {
			parent.left = next
		} else {
			parent.right = next
		}
	}
}

func findMinNode(root *Node) (*Node, *Node) {
	// 削除パターン(3)における「右部分木の最小値」と
	// その親を返す関数
	parent := root
	min := root.right
	for min.left != nil {
		parent = min
		min = min.left
	}
	return min, parent
}

func main() {
	var n int
	var root *Node
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _  = strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		operation := scanner.Text()
		switch operation {
		case "insert":
			scanner.Scan()
			value, _ := strconv.Atoi(scanner.Text())
			root = insert(root, value)
		case "print":
			printNodes(inorder(root, []int{}))
			printNodes(preorder(root, []int{}))
		case "find":
			scanner.Scan()
			target, _ := strconv.Atoi(scanner.Text())
			if find(root, target) {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		case "delete":
			scanner.Scan()
			target, _ := strconv.Atoi(scanner.Text())
			deleteNode(root, target)
		default:
			panic("No such operation!")
		}
	}
}
