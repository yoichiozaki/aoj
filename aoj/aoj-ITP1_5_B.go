
import "fmt"

func main() {
	var h, w int
	for {
		fmt.Scan(&h, &w)
		if h == 0 && w == 0 {
			break
		}
		for i := 0; i < h; i++ {
			if i == 0 || i == h - 1 {
				for j := 0; j < w; j++ {
					fmt.Print("#")
				}
				fmt.Print("\n")
				continue
			}
			fmt.Print("#")
			for l := 1; l < w - 1; l++ {
				fmt.Print(".")
			}
			fmt.Print("#")
			fmt.Print("\n")
		}
		fmt.Print("\n")
	}
}
