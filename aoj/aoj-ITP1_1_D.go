
import "fmt"

func main() {
	var x, h, m, s int
	fmt.Scan(&x)
	h = x / (60 * 60)
	m = (x - h * 60 * 60) / 60
	s = x % 60
	fmt.Printf("%d:%d:%d\n", h, m ,s)
}
