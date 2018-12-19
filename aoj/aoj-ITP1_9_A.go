
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var s string
	var ans int
	var sliced_s []string
	fmt.Scan(&s)
	for {
		scanner.Scan()
		st := scanner.Text()
		if st == "END_OF_TEXT" {
			break
		}
		st = strings.ToUpper(st)
		st = strings.ToLower(st)
		sliced_s = strings.Fields(st)
		// fmt.Println(sliced_s)
		for _, v := range sliced_s {
			if v == s {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
