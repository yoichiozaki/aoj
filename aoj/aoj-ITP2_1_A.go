
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	scanner.Scan()
	i,  e := strconv.Atoi(scanner.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	scanner.Split(bufio.ScanWords)
	arr := make([]int, 0)
	q := nextInt()
	for i := 0; i < q; i++ {
		switch nextInt() {
		case 0:
			x := nextInt()
			arr = append(arr, x)
		case 1:
			p := nextInt()
			fmt.Println(arr[p])
		case 2:
			// fmt.Printf("before: %+v\n", arr)
			arr = arr[:len(arr)-1]
			// fmt.Printf("after: %+v\n", arr)
		default:
			println("no such operation!")
		}
	}
}
