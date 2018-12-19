
import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	words := strings.Split(scanner.Text(), " ")
	wordMap := make(map[string]int)
	for _, v := range words {
		wordMap[v]++
	}
	key := ""
	value := -1
	l := -1
	word := ""
	for k, v := range wordMap {
		// fmt.Println(k, v)
		if value < v {
			key = k
			value = v
		}
		if l < len(k) {
			l = len(k)
			word = k
		}
	}

	fmt.Println(key, word)
	// fmt.Println(wordMap)
}
