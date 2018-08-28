
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"fmt"
	"math"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var bufinputs [2]float64
	var sumValue, sumCount float64
	var itemCount float64
	for scanner.Scan() {
		itemCount++
		buf := strings.Split(scanner.Text(), ",")
		bufinputs[0], _ = strconv.ParseFloat(buf[0], 64)
		bufinputs[1], _ = strconv.ParseFloat(buf[1], 64)
		sumValue += bufinputs[0]*bufinputs[1]
		sumCount += bufinputs[1]
		
	}
	fmt.Println(sumValue)
	fmt.Println(Round(sumCount / itemCount, .5, 0))
}

func Round(val float64, roundOn float64, places int) (newaVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow*val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newaVal = round / pow
	return
}
