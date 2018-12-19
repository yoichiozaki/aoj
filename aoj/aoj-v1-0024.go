
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		a, _ := strconv.ParseFloat(scanner.Text(), 64)
			t := a/9.8
			x := t*t*4.9
			ans := (x+5)/5+1
			fmt.Printf("%d\n", int(ans))
	}
}
