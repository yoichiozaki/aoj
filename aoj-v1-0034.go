
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	a := make([]float64, 10)
	for scanner.Scan() {
		info := strings.Split(scanner.Text(), ",")
		for j := 0; j < 10; j++ {
			a[j], _ = strconv.ParseFloat(info[j], 64)
		}
		v1, _ := strconv.ParseFloat(info[10], 64)
		v2, _ := strconv.ParseFloat(info[11], 64)
		sum := 0.0
		for i := 0; i < 10; i++ {
			sum += a[i]
		}
		match := v1 * ((sum/(v1+v2)))
		if match == 0 {
			fmt.Println(1)
		} else {
			sum = 0
			for k := 0; k < 10; k++ {
				sum += a[k]
				if match <= sum {
					fmt.Println(k+1)
					break
				}
			}
		}
	}
}
