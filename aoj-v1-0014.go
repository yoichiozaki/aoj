
import "fmt"

func main() {
	var d int
	for {
		v, _ := fmt.Scan(&d)
		if v == 0 {
			break
		}
		sum := 0
		for i := 0; i < 600/d; i++ {
			sum += d*d*i*d*i
		}
		fmt.Println(sum)
	}
}
