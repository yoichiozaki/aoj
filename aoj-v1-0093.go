
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	inputs := make([][]int, 0)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		a, _ := strconv.Atoi(input[0])
		b, _ := strconv.Atoi(input[1])
		if a == 0 && b == 0 {
			break
		}
		tmp := []int{a, b}
		inputs = append(inputs, tmp)
	}
	// log.Println(inputs)
	for i := 0; i < len(inputs); i++ {
		ans := leap_year(inputs[i][0], inputs[i][1])
		if len(ans) == 0 {
			fmt.Println("NA")
			fmt.Println()
		}
		for j, year := range ans {
			fmt.Println(year)
			if i != len(inputs) - 1 && j == len(ans)-1 {
				fmt.Println()
			}
		}
	}
}

func leap_year(a, b int) []int {
	ret := make([]int, 0)
	for i := a; i < b+1; i++ {
		if i % 4 == 0 && (i % 100 != 0 || i % 400 == 0) {
			ret = append(ret, i)
		}
	}
	return ret
}
