
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

var H int

func maxHeapify(A []int, i int) {
	left := 2*i
	right := 2*i + 1
	if left > H {
		return
	}
	var largest int
	if left <= H && A[left] > A[i] {
		largest = left
	} else {
		largest = i
	}
	if right <= H && A[right] > A[largest] {
		largest = right
	}

	if largest != i {
		A[i], A[largest] = A[largest], A[i]
		maxHeapify(A, largest)
	}
}

func buildMaxHeap(A []int) {
	for i := H/2; i > 0; i-- {
		maxHeapify(A, i)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	H, _ = strconv.Atoi(scanner.Text())
	heap := make([]int, H+1)
	for i := 1; i <= H; i++ {
		scanner.Scan()
		heap[i], _ = strconv.Atoi(scanner.Text())
	}
	buildMaxHeap(heap)
	for i := 1; i <= H; i++ {
		fmt.Printf(" %d", heap[i])
	}
	fmt.Println()
}
