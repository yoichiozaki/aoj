
import "fmt"

func main() {
	var s string
	var m int

	for {
		fmt.Scan(&s)
		if s == "-" {
			break
		}
		fmt.Scan(&m)
		for i := 0; i < m; i++ {
			var h int
			fmt.Scan(&h)
			s = s[h:] + s[:h]
		}
		fmt.Println(s)
	}
}
