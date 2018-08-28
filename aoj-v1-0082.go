
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	merryGoRound := []int{4,1,4,1,2,1,2,1,4,1,4,1,2,1,2,1}
	disposition := []int{41412121, 14121214, 41212141, 12121414, 21214141, 12141412, 21414121, 14141212}
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		visitors := make([]int, 0)
		for _, v := range input {
			tmp, _ := strconv.Atoi(v)
			visitors = append(visitors, tmp)
		}
		max := 0
		id := 3
		for i := 0; i < 8; i++ {
			c := 0
			for j := 0; j < 8; j++ {
				if visitors[j] > merryGoRound[i+j] {
					c += merryGoRound[i+j]
				} else {
					c += visitors[j]
				}
			}
			if max < c {
				max, id = c, i
			} else if max == c && disposition[id] > disposition[i] {
				id = i
			}
		}
		ans := strings.Split(strconv.Itoa(disposition[id]), "")
		for i := range ans {
			if i == len(ans) - 1 {
				fmt.Println(ans[i])
			} else {
				fmt.Printf("%s ", ans[i])
			}
		}
	}
}
