
import (
	"bufio"
		"os"
	"strings"
	"strconv"
		"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		ele_t1 := [3]int{}
		ele_t2 := [3]int{}
		for i := range ele_t1 {
			ele_t1[i], _ = strconv.Atoi(input[i])
			ele_t2[i], _ = strconv.Atoi(input[i+3])
		}
		if ele_t1[0] < 0 || ele_t1[1] < 0 || ele_t1[2] < 0 ||
			ele_t2[0] < 0 || ele_t2[1] < 0 || ele_t2[2] < 0 {
				break
		}
		// t1 := time.Date(ele_t1[0], time.Month(ele_t1[1]), ele_t1[2], 0, 0, 0, 0, time.UTC)
		// t2 := time.Date(ele_t2[0], time.Month(ele_t2[1]), ele_t2[2], 0, 0, 0, 0, time.UTC)
		// diff := t2.Sub(t1)
		// fmt.Printf("%v\n", int64(diff/(time.Hour*24)))
		fmt.Println(ut2jd(ele_t2[0], ele_t2[1], ele_t2[2])-ut2jd(ele_t1[0], ele_t1[1], ele_t1[2]))
	}
}

func ut2jd(year, month, day int) int {
	if month <= 2 {
		year--
		month += 12
	}
	s := 3 + year/4 - year/100 + year/400
	s += 1720994 + year*365 + (month+1)*30 + (month+1)*3/5 + day
	return s
}

