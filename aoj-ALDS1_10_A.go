
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	x, y := int64(1), int64(1)

	for i := 0; i < n; i++ {
		x, y = y, x+y
	}

	fmt.Println(x)
}
