
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
		input := strings.Split(scanner.Text(), ",")
		// log.Println(input)
		studentNo, _ := strconv.Atoi(input[0])
		weight, _ := strconv.ParseFloat(input[1], 64)
		height, _ := strconv.ParseFloat(input[2], 64)
		// log.Printf("weight = %f, height = %f\n", weight, height)
		if BMI := weight/(height*height); BMI >= 25 {
			// log.Println(BMI)
			fmt.Println(studentNo)
		}
	}
}
