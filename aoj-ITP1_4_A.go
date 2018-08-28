
import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Printf("%d %d %f\n",a / b, a % b, float64(a) / float64(b))
}
