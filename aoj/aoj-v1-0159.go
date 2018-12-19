
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"math"
	"sort"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)
const IDEAL = 22

type Person struct {
	Id int
	Weight float64
	Height float64
	BMI float64
}

type LessFunc func(p1, p2 *Person) bool
type multiSorter struct {
	People []Person
	less []LessFunc
}
func (ms *multiSorter) Sort(people []Person) {
	ms.People = people
	sort.Sort(ms)
}
func OrderBy(less ...LessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}
func (ms *multiSorter) Len() int {
	return len(ms.People)
}
func (ms *multiSorter) Swap(i, j int) {
	ms.People[i], ms.People[j] = ms.People[j], ms.People[i]
}
func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.People[i], &ms.People[j]
	var k int
	for k = 0; k < len(ms.less); k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return ms.less[k](p, q)
}

func main() {
	id := func(p1, p2 *Person) bool {
		return p1.Id < p2.Id
	}
	bmi := func(p1, p2 *Person) bool {
		return math.Abs(p1.BMI-22) < math.Abs(p2.BMI-22)
	}
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		People := make([]Person, 0)
		for i := 0; i < n; i++ {
			scanner.Scan()
			buf := strings.Split(scanner.Text(), " ")
			tmpid, _ := strconv.Atoi(buf[0])
			tmpheight, _ := strconv.ParseFloat(buf[1], 64)
			tmpweight, _ := strconv.ParseFloat(buf[2], 64)
			People = append(People,
			 		Person{tmpid,
			 		tmpheight/100,
			 		tmpweight,
			 		tmpweight/(tmpheight/100*tmpheight/100)})
		}
		OrderBy(bmi, id).Sort(People)
		// bestID := People[0].Id
		bestBMI := People[0].BMI
		for _, p := range People {
			if p.BMI == bestBMI {
				fmt.Println(p.Id)
			}
		}
	}
}
