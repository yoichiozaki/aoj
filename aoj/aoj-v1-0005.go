
import "fmt"

func GCM(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		return a
	}
	return GCM(b, a % b)
}

func LCM(a, b int) int {
	return a*b/GCM(a, b)
}

func main() {

	for {
		var a, b int
		var G, L int
		fmt.Scanf("%d %d", &a, &b)
		if a == 0 {
			break
		}
		G = GCM(a, b)
		L = LCM(a, b)
		fmt.Printf("%d %d\n", G, L)
	}
}
