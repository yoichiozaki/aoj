
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	i, e := strconv.Atoi(scanner.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func check(n int) int {
	eratos := make([]bool, n+1)
	cnt := 0
	for i := 2; i < len(eratos); i++ {
		if !eratos[i] {
			cnt++
			for j := i*i; j < len(eratos); j += i {
				eratos[j] = true
			}
		}
	}
	return cnt
}

func main() {
	for scanner.Scan() {
		n := nextInt()
		fmt.Println(check(n))
	}
}
