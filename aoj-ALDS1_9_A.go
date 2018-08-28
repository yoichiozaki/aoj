
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

const INF = 2000000001
const MAX_ELE = 250

var heapArray []int

func printParent(index int) {
	if index != 0 {
		fmt.Printf("parent key = %d, ", heapArray[(index-1)/2])
	}
}

func printLeft(parent int) {
	index := parent*2 + 1
	if index < 250 && heapArray[index] != INF {
		fmt.Printf("left key = %d, ", heapArray[index])
	}
}

func printRight(parent int) {
	index := parent*2 + 2
	if index < 250 && heapArray[index] != INF {
		fmt.Printf("right key = %d, ", heapArray[index])
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	heapArray  = make([]int, MAX_ELE)
	for i := 0; i < n; i++ {
		scanner.Scan()
		heapArray[i], _ = strconv.Atoi(scanner.Text())
	}
	for i := n; i < 250; i++ {
		heapArray[i] = INF
	}

	for i := 0; i < n; i++ {
		fmt.Printf("node %d: key = %d, ", i+1, heapArray[i])
		printParent(i)
		printLeft(i)
		printRight(i)
		fmt.Println()
	}
}
