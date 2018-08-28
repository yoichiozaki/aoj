
import "fmt"

type eachGrade struct {
	mid int
	final int
	re int
}
func main() {
	for {
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		if a == -1 && b == -1 && c == -1 {
			break
		}
		x := eachGrade{mid: a, final: b, re: c}
		if x.mid == -1 || x.final == -1{
			// fmt.Println(x)
			fmt.Println("F")
			continue
		}
		if x.mid + x.final >= 80 {
			// fmt.Println(x)
			fmt.Println("A")
			continue
		}
		if x.mid + x.final < 80 && x.mid + x.final >= 65 {
			// fmt.Println(x)
			fmt.Println("B")
			continue
		}
		if x.mid + x.final < 65 && x.mid + x.final >= 50 {
			// fmt.Println(x)
			fmt.Println("C")
			continue
		}
		if x.mid + x.final < 50 && x.mid + x.final >= 30 {
			if x.re >= 50 {
				// fmt.Println(x)
				fmt.Println("C")
				continue
			} else {
				// fmt.Println(x)
				fmt.Println("D")
				continue
			}
		}
		if x.mid + x.final < 30 {
			// fmt.Println(x)
			fmt.Println("F")
			continue
		}
	}
}
