
import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	cnt := 0
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), "")
		if len(buf) == 1 {
			cnt++
			continue
		}
		flag := true
		for i := 0; i < len(buf)/2; i++ {
			if buf[i] != buf[len(buf)-1-i] {
				flag = false
				break
			}
		}
		if flag {
			cnt++
		}
	}
	fmt.Println(cnt)
}
