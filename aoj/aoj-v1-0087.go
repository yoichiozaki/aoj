
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
		stack := make([]float64, 0)
		for _, v := range input {
			switch v {
			case "+":
				y := stack[len(stack)-1]
				x := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, x+y)
			case "-":
				y := stack[len(stack)-1]
				x := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, x-y)
			case "/":
				y := stack[len(stack)-1]
				x := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, x/y)
			case "*":
				y := stack[len(stack)-1]
				x := stack[len(stack)-2]
				stack = stack[:len(stack)-2]
				stack = append(stack, x*y)
			default:
				tmp, _ := strconv.ParseFloat(v, 64)
				stack = append(stack, tmp)
			}
		}
		fmt.Printf("%f\n", stack[0])
	}
}
