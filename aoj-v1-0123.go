
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	m500 := []float64{35.5, 37.5, 40.0, 43.0, 50.0, 55.0, 70.0, 1000.0}
	m1000 := []float64{71.0, 77.0, 83.0, 89.0, 105.0, 116.0, 148.0, 1000.0}
	classes := []string{"AAA", "AA", "A", "B", "C", "D", "E", "NA"}
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		t500, _ := strconv.ParseFloat(input[0], 64)
		t1000, _ := strconv.ParseFloat(input[1], 64)
		for i := 0; i < len(m500); i++ {
			if t500 < m500[i] && t1000 < m1000[i] {
				fmt.Println(classes[i])
				break
			}
		}
	}
}
