
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	scanner.Scan()
	i, e := strconv.Atoi(scanner.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func binarySearch(target int, ints []int) bool {
	low := 0
	high := len(ints)-1
	for low <= high {
		mid := (low + high)/2
		guess := ints[mid]
		if guess == target {
			return true
		}
		if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

func main() {
	scanner.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]int, n)
	for i :=range a {
		a[i] = nextInt()
	}
	q := nextInt()
	for i :=0; i < q; i++ {
		k := nextInt()
		if binarySearch(k, a) {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}
