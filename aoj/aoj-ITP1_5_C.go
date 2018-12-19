
import "fmt"

func main() {
	var h, w int
	for {
		fmt.Scan(&h, &w)
		if h == 0 && w == 0 {
			break
		}
		for i := 0; i < h; i++ {
			if i % 2 == 0 {
				for j := 0; j < w; j++ {
					if j != w - 1 {
						if j % 2 == 0 {
							fmt.Print("#")
						} else {
							fmt.Print(".")
						}
						continue
					}
					if w % 2 == 0 {
						fmt.Print(".\n")
					} else {
						fmt.Print("#\n")
					}
				}
				continue
			} else {
				for j := 0; j < w; j++ {
					if j != w - 1 {
						if j % 2 == 0 {
							fmt.Print(".")
						} else {
							fmt.Print("#")
						}
						continue
					}
					if w % 2 == 0 {
						fmt.Print("#\n")
					} else {
						fmt.Print(".\n")
					}
				}
			}
		}
		fmt.Print("\n")
	}
}
