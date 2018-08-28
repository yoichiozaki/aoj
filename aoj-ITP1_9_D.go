
import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"strconv"
)

func reverse(s string) string{
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1{
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var(
		str, p string
		n, a, b int
	)
	fmt.Scan(&str, &n)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		x := scanner.Text()
		xl := strings.Fields(x)
		// fmt.Println(xl)
		switch xl[0] {
		case "replace":
			a, _ = strconv.Atoi(xl[1])
			b, _ = strconv.Atoi(xl[2])
			p = xl[3]
			// fmt.Println(p)
			str = str[:a] + p + str[b+1:]
		case "reverse":
			a, _ = strconv.Atoi(xl[1])
			b, _ = strconv.Atoi(xl[2])
			str = str[:a] + reverse(str[a:b+1]) + str[b+1:]
			// fmt.Println(str)
		case "print":
			a, _ = strconv.Atoi(xl[1])
			b, _ = strconv.Atoi(xl[2])
			// fmt.Println(str)
			fmt.Println(str[a:b+1])
		}
	}
}
