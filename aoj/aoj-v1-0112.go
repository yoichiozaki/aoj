
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		f := make([]int, 65)
		for i := 0; i < n; i++ {
			scanner.Scan()
			t, _ := strconv.Atoi(scanner.Text())
			f[t]++
		}
		if n == 1 {
			fmt.Println(0)
			continue
		}
		i, ans, m := 1, 0, n-f[0]
		for m > 0 {
			k := f[i]
			if k == 0 {
				i++
				continue
			}
			if k == 1 {
				ans += (m-1)*i
			} else {
				a := m*k - (k*(k+1))/2
				ans += a*i
			}
			m -= k
			i++
		}
		fmt.Println(ans)
	}
}

