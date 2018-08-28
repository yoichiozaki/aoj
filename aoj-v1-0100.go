
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		achievement := make([]int, 4000)
		members := make([]int, 0)
		check := func(target int) bool {
			for _, v := range members {
				if v == target {
					return true
				}
			}
			return false
		}
		for i := 0; i < n; i++ {
			scanner.Scan()
			input := strings.Split(scanner.Text(), " ")
			e, _ := strconv.Atoi(input[0])
			p, _ := strconv.Atoi(input[1])
			q, _ := strconv.Atoi(input[2])
			achievement[e] += p*q
			if !check(e) {
				members = append(members, e)
			}
		}
		// log.Println(achievement)
		flag := false
		for _, v := range members {
			if achievement[v] >= 1000000 {
				fmt.Println(v)
				flag = true
			}
		}
		if !flag {
			fmt.Println("NA")
		}
	}
}
