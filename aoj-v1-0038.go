
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), ",")
		cards := make([]int, 5)
		for i :=range input {
			cards[i], _ = strconv.Atoi(input[i])
		}
		sort.Sort(sort.IntSlice(cards))
		if cards[0] == cards[1] && cards[2] == cards[3] && cards[2] == cards[1] {
			fmt.Println("four card")
		} else if cards[1] == cards[2] && cards[3] == cards[4] && cards[3] == cards[2]{
			fmt.Println("four card")
		} else if cards[0] == cards[1] && cards[1] == cards[2] && cards[3] == cards[4] {
			fmt.Println("full house")
		} else if cards[0] == cards[1] && cards[2] == cards[3] && cards[3] == cards[4] {
			fmt.Println("full house")
		} else if cards[0] == 1 && cards[1] == 10 && cards[2] == 11 && cards[3] == 12 && cards[4] == 13 {
			fmt.Println("straight")
		} else if cards[0]+1 == cards[1] && cards[1]+1 == cards[2] && cards[2]+1 == cards[3] && cards[3]+1 == cards[4] {
			fmt.Println("straight")
		} else if (cards[0] == cards[1] && cards[1] == cards[2]) || (cards[1] == cards[2] && cards[2] == cards[3]) || (cards[2] == cards[3] && cards[3] == cards[4]) {
			fmt.Println("three card")
		} else if ((cards[0] == cards[1]) && ((cards[2] == cards[3]) || (cards[3] == cards[4]))) || ((cards[1] == cards[2]) && (cards[3] == cards[4])) {
			fmt.Println("two pair")
		} else if cards[0] == cards[1] || cards[1] == cards[2] || cards[2] == cards[3] || cards[3] == cards[4] {
			fmt.Println("one pair")
		} else {
			fmt.Println("null")
		}
	}
}
