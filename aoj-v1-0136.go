
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	result := make([]int, 6)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		h, _ := strconv.ParseFloat(scanner.Text(), 64)
		if h < 165.0 {
			result[0]++
		} else if 165.0 <= h && h < 170.0 {
			result[1]++
		} else if 170.0 <= h && h < 175.0 {
			result[2]++
		} else if 175.0 <= h && h < 180.0 {
			result[3]++
		} else if 180.0 <= h && h < 185.0 {
			result[4]++
		} else if 185.0 <= h {
			result[5]++
		}
	}
	for i := 0; i < 6; i++ {
		fmt.Printf("%d:", i+1)
		for j := 0; j < result[i]; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
