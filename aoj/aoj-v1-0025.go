
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	a := make([]int, 4)
	b := make([]int ,4)
	for scanner.Scan() {
		inputa := strings.Split(scanner.Text(), " ")
		for i := range a {
			a[i], _ = strconv.Atoi(inputa[i])
		}
		scanner.Scan()
		inputb := strings.Split(scanner.Text(), " ")
		for i := range b {
			b[i], _ = strconv.Atoi(inputb[i])
		}
		hit := 0
		blow := 0
		if a[0] == b[0] {
			hit++
		}
		if a[1] == b[1] {
			hit++
		}
		if a[2] == b[2] {
			hit++
		}
		if a[3] == b[3] {
			hit++
		}
		if a[0] == b[1] {
			blow++
		}
		if a[1] == b[2] {
			blow++
		}
		if a[2] == b[3] {
			blow++
		}
		if a[1] == b[0] {
			blow++
		}
		if a[2] == b[1] {
			blow++
		}
		if a[2] == b[0] {
			blow++
		}
		if a[3] == b[2] {
			blow++
		}
		if a[3] == b[1] {
			blow++
		}
		if a[3] == b[0] {
			blow++
		}
		if a[0] == b[2] {
			blow++
		}
		if a[1] == b[3] {
			blow++
		}
		if a[0] == b[3] {
			blow++
		}
		fmt.Printf("%d %d\n", hit, blow)
	}
}
