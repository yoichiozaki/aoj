
import "fmt"

func main() {
	var i, counter int
	counter = 1
	for {
		fmt.Scan(&i)
		if i == 0 {
			break
		}
		fmt.Printf("Case %d: %d\n", counter, i)
		counter++
	}
}
