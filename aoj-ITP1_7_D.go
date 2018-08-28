
import "fmt"

func main() {
	var n, m, l int
	fmt.Scan(&n, &m, &l)
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			fmt.Scan(&a[i][j])
		}
	}
	b := make([][]int, m)
	for i := range b {
		b[i] = make([]int, l)
		for j := range b[i] {
			fmt.Scan(&b[i][j])
		}
	}
	// fmt.Println(a)
	// fmt.Println(b)

	c := make([][]int, n)
	for i := range c {
		c[i] = make([]int, l)
	}
	// fmt.Println(c)

	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			for k := 0; k < m; k++ {
				c[i][j] += a[i][k]*b[k][j]
			}
		}
	}
	// fmt.Println(c)
	for i := 0; i< n; i++ {
		for j := 0; j < l; j++ {
			if j == l - 1 {
				fmt.Printf("%d\n", c[i][j])
			} else {
				fmt.Printf("%d ", c[i][j])
			}
		}
	}
}
