
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		weight, _ := strconv.ParseFloat(scanner.Text(), 64)
		if weight <= 48 {
			fmt.Println("light fly")
			continue
		}
		if weight > 48 && weight <= 51 {
			fmt.Println("fly")
			continue
		}
		if weight > 51 && weight <= 54 {
			fmt.Println("bantam")
			continue
		}
		if weight > 54 && weight <= 57 {
			fmt.Println("feather")
			continue
		}
		if weight > 57 && weight <= 60 {
			fmt.Println("light")
			continue
		}
		if weight > 60 && weight <= 64 {
			fmt.Println("light welter")
			continue
		}
		if weight > 64 && weight <= 69 {
			fmt.Println("welter")
			continue
		}
		if weight > 69 && weight <= 75 {
			fmt.Println("light middle")
			continue
		}
		if weight > 75 && weight <= 81 {
			fmt.Println("middle")
			continue
		}
		if weight > 81 && weight <= 91 {
			fmt.Println("light heavy")
			continue
		}
		if weight > 91 {
			fmt.Println("heavy")
			continue
		}
	}
}
