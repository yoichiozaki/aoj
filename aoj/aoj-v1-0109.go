
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"log"
	"go/token"
	"go/types"
	"go/constant"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), "")
		src := input[:len(input)-1]
		// log.Println(src)
		fs := token.NewFileSet()
		tv, err := types.Eval(fs, nil, token.NoPos, strings.Join(src, ""))
		if err != nil {
			log.Println(err)
			return
		}
		val, ok := constant.Int64Val(tv.Value)
		if !ok {
			log.Println("Failed to get Int64Val")
			return
		}
		fmt.Println(val)
	}
}
