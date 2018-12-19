
import "fmt"

func main() {
	var a, b, c, counter int
	fmt.Scan(&a, &b, &c)
	for i := a; i <= b; i++ {
		if c % i == 0 {
			counter++
		}
	}
	fmt.Println(counter)
}
