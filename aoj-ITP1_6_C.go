
import (
	"fmt"
	"strings"
)

func main() {
	var n, b, f, r ,v int
	fmt.Scan(&n)
	building1 := make([][]int, 3)
	building2 := make([][]int, 3)
	building3 := make([][]int, 3)
	building4 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		building1[i] = make([]int, 10)
		building2[i] = make([]int, 10)
		building3[i] = make([]int, 10)
		building4[i] = make([]int, 10)
	}
	// fmt.Println(building1)


	for i := 0; i < n; i++ {
		fmt.Scan(&b, &f, &r, &v)
		switch b {
		case 1: building1[f-1][r-1] += v
		case 2: building2[f-1][r-1] += v
		case 3: building3[f-1][r-1] += v
		case 4: building4[f-1][r-1] += v
		}
	}

	for i := 0; i < 3; i ++ {
		for j := 0; j < 10; j++ {
			fmt.Printf(" %d", building1[i][j])
		}
		fmt.Println()

	}
	fmt.Println(strings.Repeat("#", 20))
	for i := 0; i < 3; i ++ {
		for j := 0; j < 10; j++ {
			fmt.Printf(" %d", building2[i][j])
		}
		fmt.Println()

	}
	fmt.Println(strings.Repeat("#", 20))
	for i := 0; i < 3; i ++ {
		for j := 0; j < 10; j++ {
			fmt.Printf(" %d", building3[i][j])
		}
		fmt.Println()

	}
	fmt.Println(strings.Repeat("#", 20))
	for i := 0; i < 3; i ++ {
		for j := 0; j < 10; j++ {
			fmt.Printf(" %d", building4[i][j])
		}
		fmt.Println()

	}
}
