
import (
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Scan(&r)
	fmt.Printf("%f %f\n", math.Pi*float64(r)*float64(r), math.Pi*float64(2 * r))
}
