
import "fmt"

func main() {
	var (
		a, b, ans int
		op string
	)
	for {
		fmt.Scan(&a, &op, &b)
		if op == "?" {
			break
		}
		switch op {
		case "+": ans = a + b
		case "-": ans = a - b
		case "*": ans = a * b
		case "/": ans = a / b
		default:
			fmt.Println("op is incorrect.")
		}
		fmt.Println(ans)
	}
}
