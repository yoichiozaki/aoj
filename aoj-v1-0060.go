
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

var scanner  = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		inputs := strings.Split(scanner.Text(), " ")
		var check [10]bool
		for i := 0; i < 10; i++ {
			check[i] = false
		}
		a, _ := strconv.Atoi(inputs[0])
		b, _ := strconv.Atoi(inputs[1])
		c, _ := strconv.Atoi(inputs[2])
		check[a-1] = true
		check[b-1] = true
		check[c-1] = true
		cnt := 0
		for i := 0; i < 10; i++ {
			if !check[i] {
				if a+b+1+i <= 20 {
					cnt++
				}
			}
		}
		if cnt > 3 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		cnt = 0
	}
}
