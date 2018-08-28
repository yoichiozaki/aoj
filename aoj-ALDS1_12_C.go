
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

type Node struct {
	id int
	edgesTo []*Node
	edgesCost []int
	isVisited bool
	cost int
	previousNode *Node
}

func (n *Node) connect(node *Node, cost int) {
	n.edgesTo = append(n.edgesTo, node)
	n.edgesCost = append(n.edgesCost, cost)
}

var scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	scanner.Scan()
	i, e := strconv.Atoi(scanner.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	scanner.Split(bufio.ScanWords)
	n := nextInt()
	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = &Node{cost:-1}
	}
	for i := 0; i < n; i++ {
		u := nextInt()
		k := nextInt()
		nodes[i].id = u
		for j := 0; j < k;j++ {
			v := nextInt()
			c := nextInt()
			nodes[i].connect(nodes[v], c)
		}
	}

	nodes[0].cost = 0
	for {
		var processing *Node
		processing = nil
		for i := 0; i < len(nodes); i++ {
			node := nodes[i]
			if node.isVisited || node.cost < 0 {
				continue
			}
			if processing == nil {
				processing = node
				continue
			}
			if node.cost < processing.cost {
				processing = node
			}
		}
		if processing == nil {
			break
		}
		processing.isVisited = true
		for i := 0; i < len(processing.edgesTo); i++ {
			node := processing.edgesTo[i]
			cost := processing.cost + processing.edgesCost[i]
			if node.cost < 0 || node.cost > cost {
				node.cost = cost
				node.previousNode = processing
			}
		}
	}

	for _, n := range nodes {
		fmt.Printf("%d %d\n",n.id, n.cost)
	}

}
