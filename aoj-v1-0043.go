
import (
	"bufio"
	"os"
	"fmt"
)

func isOK(new_table [10]int) bool {
	trio := 0
	for k := 1; k <= 7; k++ {
		if new_table[k] > 2 {
			trio++
			new_table[k] -= 3
		}
		for new_table[k] > 0 && new_table[k+1] > 0 && new_table[k+2] > 0 {
			trio++
			new_table[k]--
			new_table[k+1]--
			new_table[k+2]--
		}
	}
	for k := 1; k<= 9; k++ {
		if new_table[k] == 3 {
			trio++
			new_table[k] -= 3
		}
	}
	if trio == 4 {
		return true
	}
	return false
}
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		var table [10]int
		buffer := scanner.Text()
		for _, c := range buffer{
			i := int(c - '0')
			table[i]++
		}
		First := true
		for i := 1; i <= 9; i++ {
			if table[i] < 4 {
				var new_table [10]int
				for k := 1; k <= 9; k++ {
					new_table[k] = table[k]
				}
				new_table[i]++
				for k := 1; k <= 9; k++ {
					if new_table[k] >= 2 {
						var new_table2 [10]int
						for p := 1; p <= 9; p++ {
							new_table2[p] = new_table[p]
						}
						new_table2[k] -= 2
						if isOK(new_table2) {
							if First {
								fmt.Printf("%d", i)
								First = false
							} else {
								fmt.Printf(" %d", i)
							}
							break
						}
					}
				}
			}
		}
		if First {
			fmt.Printf("0\n")
		} else {
			fmt.Printf("\n")
		}
	}
}

