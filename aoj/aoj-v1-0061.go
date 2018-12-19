
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

type team struct {
	id, num, rank int
}

func (x *team) initialize() {
	x.id = -1
	x.rank = 0
	x.num = -1
}

func main() {
	var result [101]int
	for i, _ := range result {
		result[i] = 0
	}
	teams := make([]team, 0)
	for i := 0; i < 101; i++ {
		var tmp team
		tmp.initialize()
		teams = append(teams, tmp)
	}
	cnt := 0
	for scanner.Scan() {
		cnt++
		input := strings.Split(scanner.Text(), ",")
		if input[0] == "0" && input[1] == "0" {
			break
		}
		id, _ := strconv.Atoi(input[0])
		number, _ := strconv.Atoi(input[1])
		teams[id].num = number
		teams[id].id = id
	}
	cnt--
	// for i := 0; i < 101; i++ {
	// 	fmt.Printf("%+v\n", teams[i])
	// }
	for i := 1; i <= 100; i++ {
		for k := 100; k >= i; k-- {
			if teams[k].num > teams[k-1].num {
				teams[k], teams[k-1] = teams[k-1], teams[k]
			}
		}
	}
	// fmt.Println("----------")
	// for i := 0; i < cnt; i++ {
	// 	fmt.Printf("%+v\n", teams[i])
	// }
	// os.Exit(1)
	rank := 0
	tmpnum := 0
	for i := 0; i < cnt; i++ {
		if teams[i].num != tmpnum {
			rank++
			teams[i].rank = rank
			result[teams[i].id] = rank
			tmpnum = teams[i].num
		} else {
			teams[i].rank = rank
			result[teams[i].id] = rank
			tmpnum = teams[i].num
		}
	}
	for scanner.Scan() {
		q, _ := strconv.Atoi(scanner.Text())
		fmt.Printf("%d\n", result[q])
	}
}
