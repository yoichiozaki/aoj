
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		// fmt.Println(input)
		points := make([]float64, 0)
		for _, x := range input {
			p, _ := strconv.ParseFloat(x, 64)
			points = append(points, p)
		}
		if points[0] > points[6] || points[3] < points[5] || points[4] > points[2] || points[7] < points[1] {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
		points = points[:1]
	}
}
