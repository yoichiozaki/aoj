
import (
	"fmt"
	"os"
	"math"
)

func getAverage(nums []float64) float64{
	var sum float64
	for _, v := range nums {
		sum += v
	}
	return sum/float64(len(nums))
}

func main() {
	for {
		var n int
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		scores := make([]float64, n)
		for i := range scores {
			fmt.Scan(&scores[i])
		}
		// fmt.Println(scores)
		aveTwo := math.Pow(getAverage(scores), 2)
		for i, v := range scores {
			scores[i] = math.Pow(v, 2)
		}
		// fmt.Println(scores)
		twoAve := getAverage(scores)
		fmt.Printf("%f\n", math.Sqrt(twoAve - aveTwo))
	}
	os.Exit(0)
}

