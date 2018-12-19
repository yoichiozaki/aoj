
import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			fmt.Scan(&a[i][j])
		}
	}
	// fmt.Println(a)
	b := make([]int, m)
	for i := range b {
		fmt.Scan(&b[i])
	}
	// fmt.Println(b)
	ans := make([]int, n)
	for k := 0; k < n; k++ {
		for i := 0; i < m; i ++ {
			// fmt.Printf("a[k][i]: %d, b[i]: %d\n", a[k][i], b[i])
			ans[k] += a[k][i] * b[i]
		}
	}

	for _, v := range ans {
		fmt.Println(v)
	}
}
