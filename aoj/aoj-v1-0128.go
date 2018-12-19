
import (
	"bufio"
	"fmt"
		"os"
	"strconv"
	"strings"
	)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	table := [10]string{
		"* = ****",
		"* =* ***",
		"* =** **",
		"* =*** *",
		"* =**** ",
		" *= ****",
		" *=* ***",
		" *=** **",
		" *=*** *",
		" *=**** ",
	}
	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	n := len(input)
	cnt := 0
	for cnt < n {
		tmp := strings.Split(fmt.Sprintf("%05s", input[cnt]), "")
		ans := make([]string, 0)
		for _, v := range tmp {
			iv, _ := strconv.Atoi(v)
			ans = append(ans, table[iv])
		}
		for i := 0; i < 8; i++ {
			for j := 0; j < 5; j++ {
				fmt.Print(string(ans[j][i]))
			}
			fmt.Println()
		}
		if cnt != n-1 {
			fmt.Println()
		}
		cnt++
	}
}

