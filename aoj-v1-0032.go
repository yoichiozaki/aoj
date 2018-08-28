
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	d := 0
	e := 0
	for scanner.Scan() {
		info := strings.Split(scanner.Text(), ",")
		a, _ := strconv.Atoi(info[0])
		b, _ := strconv.Atoi(info[1])
		c, _ := strconv.Atoi(info[2])
		if a*a + b*b == c*c {
			d++
		}
		if a == b {
			e++
		}
	}
	fmt.Println(d)
	fmt.Println(e)
}
