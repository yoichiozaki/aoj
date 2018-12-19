
import "fmt"

func main() {
	var r, c int
	fmt.Scan(&r, &c)
	sheet := make([][]int, r)
	for i := range sheet {
		sheet[i] = make([]int, c)
		for j := 0; j < c; j++ {
			fmt.Scan(&sheet[i][j])
		}
	}
	// fmt.Println(sheet)
	for i := 0; i < r; i++ {
		x := 0
		for j := 0; j < c; j++ {
			x += sheet[i][j]
		}
		sheet[i] = append(sheet[i], x)
	}
	sheet = append(sheet, make([]int, c+1))
	for i := 0; i <= c; i++ {
		y := 0
		for j := 0; j < r; j++ {
			y += sheet[j][i]
		}
		sheet[r][i] = y
	}
	// fmt.Println(sheet)
	for i := 0; i < r+1; i++{
		for j := 0; j < c+1; j++ {
			if j == c {
				fmt.Printf("%d\n", sheet[i][j])
			} else {
				fmt.Printf("%d ", sheet[i][j])
			}
		}
	}
}
