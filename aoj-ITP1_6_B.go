
import "fmt"

func main(){
	var n, num int
	var s string
	fmt.Scan(&n)
	inputCards := make([][]bool, 4)
	for i := 0; i < 4; i++ {
		inputCards[i] = make([]bool, 13)
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&s, &num)
		// fmt.Println(s, num)
		switch s{
		case "S": inputCards[0][num-1] = true
		case "H": inputCards[1][num-1] = true
		case "C": inputCards[2][num-1] = true
		case "D": inputCards[3][num-1] = true
		default: fmt.Println("input is invalid.")
		}
	}
	// fmt.Println(inputCards)

	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			switch i {
			case 0:
				if inputCards[i][j] == false {
					fmt.Printf("S %d\n", j+1)
				}
			case 1:
				if inputCards[i][j] == false {
					fmt.Printf("H %d\n", j+1)
				}
			case 2:
				if inputCards[i][j] == false {
					fmt.Printf("C %d\n", j+1)
				}
			case 3:
				if inputCards[i][j] == false {
					fmt.Printf("D %d\n", j+1)
				}
			}
		}
	}
}
