
import "fmt"

func main() {
	var i int
	stack := make([]int, 0)
	for {
		v, _ := fmt.Scan(&i)
		// fmt.Println(stack)
		// println(v)
		if v == 0 {
			break
		}
		if i == 0 {
			fmt.Println(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, i)
	}
}
