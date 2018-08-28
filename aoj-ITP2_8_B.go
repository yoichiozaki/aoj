
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

func nextString() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	scanner.Split(bufio.ScanWords)
	dict := make(map[string]int)
	q := nextInt()
	for i := 0; i < q; i++ {
		switch nextInt() {
		case 0:
			key := nextString()
			x := nextInt()
			if _, ok := dict[key]; ok {
				dict[key] = x
			} else {
				dict[key] = x
			}
		case 1:
			key := nextString()
			if _, ok := dict[key]; ok {
				fmt.Println(dict[key])
			} else {
				fmt.Println(0)
			}
		case 2:
			key := nextString()
			delete(dict, key)
		default:
			panic("no such operation!")
		}
	}
}
