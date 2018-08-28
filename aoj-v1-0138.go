
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	race1 := make([]player, 8)
	race2 := make([]player, 8)
	race3 := make([]player, 8)
	rest := make([]player, 6)
	for i := 0; i < 8; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		no, _ := strconv.Atoi(input[0])
		time, _ := strconv.ParseFloat(input[1], 64)
		race1[i] = player{no, time}
	}
	for i := 0; i < 8; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		no, _ := strconv.Atoi(input[0])
		time, _ := strconv.ParseFloat(input[1], 64)
		race2[i] = player{no, time}
	}
	for i := 0; i < 8; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		no, _ := strconv.Atoi(input[0])
		time, _ := strconv.ParseFloat(input[1], 64)
		race3[i] = player{no, time}
	}
	for i := 1; i <= 7; i++ {
		for k := 7; k >= i; k-- {
			if race1[k].time < race1[k-1].time {
				race1[k], race1[k-1] = race1[k-1], race1[k]
			}
			if race2[k].time < race2[k-1].time {
				race2[k], race2[k-1] = race2[k-1], race2[k]
			}
			if race3[k].time < race3[k-1].time {
				race3[k], race3[k-1] = race3[k-1], race3[k]
			}
		}
	}
	rest[0] = race1[2]
	rest[1] = race1[3]
	rest[2] = race2[2]
	rest[3] = race2[3]
	rest[4] = race3[2]
	rest[5] = race3[3]
	for i := 1; i <= 5; i++ {
		for k := 5; k >= i; k-- {
			if rest[k].time < rest[k-1].time {
				rest[k],rest[k-1] = rest[k-1], rest[k]
			}
		}
	}
	fmt.Printf("%d %.2f\n", race1[0].no, race1[0].time)
	fmt.Printf("%d %.2f\n", race1[1].no, race1[1].time)
	fmt.Printf("%d %.2f\n", race2[0].no, race2[0].time)
	fmt.Printf("%d %.2f\n", race2[1].no, race2[1].time)
	fmt.Printf("%d %.2f\n", race3[0].no, race3[0].time)
	fmt.Printf("%d %.2f\n", race3[1].no, race3[1].time)
	fmt.Printf("%d %.2f\n", rest[0].no, rest[0].time)
	fmt.Printf("%d %.2f\n", rest[1].no, rest[1].time)
}

type player struct {
	no int
	time float64
}
