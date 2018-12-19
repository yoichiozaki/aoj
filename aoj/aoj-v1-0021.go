
import "fmt"

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n int
	var x1, y1, x2, y2, x3, y3, x4, y4 float64
	fmt.Scan(&n)
	for i := 0 ; i < n; i++ {
		fmt.Scanf("%f %f %f %f %f %f %f %f", &x1, &y1, &x2, &y2, &x3, &y3, &x4, &y4)
		if abs((x2-x1)*(y4-y3) - (y2-y1)*(x4-x3)) < 1e-10 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
