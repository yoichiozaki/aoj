
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func nextInt() int {
	scanner.Scan()
	i, e := strconv.Atoi(scanner.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func zeller(month, day int) int {
	year := 2004
	if month == 1 || month == 2 {
		year--
		month += 12
	}
	return (year + year/4 - year/100 + year/400 + (13*month+8)/5 + day)%7
}

func main() {
	scanner.Split(bufio.ScanWords)
	for {
		month := nextInt()
		day := nextInt()
		if month == 0 && day == 0{
			break
		}
		switch zeller(month, day) {
		case 0:
			fmt.Println("Sunday")
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
		}
	}
}
