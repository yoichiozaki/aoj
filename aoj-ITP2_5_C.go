
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"reflect"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	scanner.Scan()
	i, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func isMember(target int, ints []int) bool {
	for i := range ints {
		if ints[i] == target {
			return true
		}
	}
	return false
}

func _perm(f func([]int), n, m int, xs []int) {
	if len(xs) == m {
		f(xs)
	} else {
		for i := 0; i <= n; i++ {
			if !isMember(i, xs) {
				_perm(f, n, m, append(xs, i))
			}
		}
	}
}

func permutation(f func([]int), n, m int) {
	_perm(f, n, m, make([]int, 0, m))
}

func printArray(target []int) {
	for i := range target {
		if i == len(target)-1 {
			fmt.Printf("%d\n", target[i])
		} else {
			fmt.Printf("%d ", target[i])
		}
	}
}

func main() {
	scanner.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]int, n)
	for i := range a {
		a[i] = nextInt()
	}
	store := make([][]int, 0)
	// fmt.Println(store)
	var i int
	printPermutation := func(xs []int) {
		// fmt.Print(xs)
		// fmt.Println()
		store = append(store, make([]int, 0))
		store[i] = append(store[i], xs...)
		// fmt.Println(store)
		i++
	}
	permutation(printPermutation, n-1, n)
	// fmt.Println(store)
	for i := range store {
		for j := range store[i] {
			store[i][j]++
		}
	}
	// fmt.Println(store)
	for i := range store {
		if reflect.DeepEqual(store[i], a) {
			if i-1 >= 0 {
				printArray(store[i-1])
			}
			printArray(store[i])
			if i+1 < len(store) {
				printArray(store[i+1])
			}
		}
	}
}
