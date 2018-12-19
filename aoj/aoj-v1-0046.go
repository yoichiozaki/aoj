
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
	"sort"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	heights := make([]float64, 0)
	for scanner.Scan() {
		height, _ := strconv.ParseFloat(scanner.Text(), 64)
		heights = append(heights, height)
		// fmt.Println(heights)
	}
	sort.Float64s(heights)
	fmt.Println(heights[len(heights)-1] - heights[0])
}
