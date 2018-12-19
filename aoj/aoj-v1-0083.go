
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
		time := make([]int, 0)
		for _, v := range input {
			tmp, _ := strconv.Atoi(v)
			time = append(time, tmp)
		}
		if time[0] < 1868 || time[0] == 1868 && (time[1] < 9 || (time[1] == 9 && time[2] < 8)) {
			fmt.Println("pre-meiji")
		} else if time[0] < 1912 || time[0] == 1912 && (time[1] < 7 || (time[1] == 7 && time[2] < 30)) {
			fmt.Printf("meiji %d %d %d\n", time[0] - 1867, time[1], time[2])
		} else if time[0] < 1926 || time[0] == 1926 && (time[1] < 12 || (time[1] == 12 && time[2] < 25)) {
			fmt.Printf("taisho %d %d %d\n", time[0] - 1911, time[1], time[2])
		} else if time[0] < 1989 || time[0] == 1989 && (time[1] < 1 || (time[1] == 1 && time[2] < 8)) {
			fmt.Printf("showa %d %d %d\n", time[0] - 1925, time[1], time[2])
		} else {
			fmt.Printf("heisei %d %d %d\n", time[0] - 1988, time[1], time[2])
		}
	}
}
