
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
	dict := make(map[string][]int)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		word := input[0]
		index, _ := strconv.Atoi(input[1])
		if _, ok := dict[word]; ok {
			dict[word] = append(dict[word], index)
		} else {
			dict[word] = []int{index}
		}
		// log.Println(dict)
	}
	keys := make([]string, len(dict))
	i := 0
	for key, _ := range dict {
		keys[i] = key
		i++
	}
	sort.Strings(keys)
	for i := range keys {
		if v, ok := dict[keys[i]]; ok {
			sort.Ints(v)
			fmt.Println(keys[i])
			for j := range v {
				if j == len(v)-1 {
					fmt.Printf("%d\n", v[j])
				} else {
					fmt.Printf("%d ", v[j])
				}
			}
		}
	}
}
