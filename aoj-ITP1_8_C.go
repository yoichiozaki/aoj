
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	alphabetsMap := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0, "h": 0, "i": 0, "j": 0,
								"k": 0, "l": 0, "m": 0, "n": 0, "o": 0, "p": 0, "q": 0, "r": 0, "s": 0, "t": 0,
								"u": 0, "v": 0, "w": 0, "x": 0, "y": 0, "z": 0}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		st := scanner.Text()
		if st == "" {
			break
		}
		sr := ""
		for _, v := range st {
			x := string(v)
			l := strings.ToLower(x)
			if l != x {
				sr += l
			} else {
				sr += x
			}
		}
		// fmt.Println(sr)
		for _, v := range sr {
			alphabetsMap[string(v)]++
		}
		// fmt.Println(alphabetsMap)
	}
	for i := 0; i < len(alphabet); i++ {
		fmt.Printf("%s : %d\n", alphabet[i:i+1], alphabetsMap[alphabet[i:i+1]])
	}
}

