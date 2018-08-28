
import (
	"fmt"
	"math"
)

func abs(x float64) float64 {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func getSquaring(x float64) float64 {
	return x*x
}

func getQube(x float64) float64 {
	return x*x*x
}

func main() {
	var n, ans float64
	fmt.Scan(&n)
	n_int := int(n)
	x := make([]float64, n_int)
	y := make([]float64, n_int)
	for i := range x {
		fmt.Scan(&x[i])
	}
	for i := range y {
		fmt.Scan(&y[i])
	}

	// p == 1
	ans = 0
	for i := 0; i < n_int; i++ {
		ans += abs(x[i] - y[i])
	}
	fmt.Printf("%f\n", ans)

	// p == 2
	ans = 0
	for i := 0; i < n_int; i++ {
		ans += getSquaring(abs(x[i] - y[i]))
	}
	fmt.Println(math.Sqrt(ans))

	// p == 3
	ans = 0
	for i := 0; i < n_int; i++ {
		ans += getQube(abs(x[i] - y[i]))
	}
	fmt.Println(math.Cbrt(ans))

	// p == +Inf
	ans = -1
	for i := 0; i < n_int; i++ {
		ans = math.Max(ans, abs(x[i] - y[i]))
	}
	fmt.Printf("%f\n", ans)
}
