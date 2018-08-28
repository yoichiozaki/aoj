
import (
	"container/heap"
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	h := &IntHeap{}
	heap.Init(h)

	LOOP:
	for {
		scanner.Scan()
		switch scanner.Text() {
		case "end":
			break LOOP
		case "insert":
			scanner.Scan()
			element, _ := strconv.Atoi(scanner.Text())
			heap.Push(h, element)
		case "extract":
			fmt.Printf("%d\n", heap.Pop(h))
		default:
			panic("No such operation.")
		}
	}
}
