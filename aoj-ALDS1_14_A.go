
import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	scanner :=bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	scanner.Scan()
	target := scanner.Text()
	n := len(input)
	m := len(target)
	ans := make([]int, 0)
	var p int
	for i := 0; i <= n-m; i++ {
		for p = 0; p < m; p++ {
			if input[i+p] != target[p] {
				break
			}
		}
		if p == m {
			ans = append(ans, i)
		}
	}
	for i :=range ans{
		fmt.Println(ans[i])
	}
}
