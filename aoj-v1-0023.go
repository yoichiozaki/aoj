
import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var xa, ya, ra, xb, yb, rb float64
		fmt.Scan(&xa, &ya, &ra, &xb, &yb, &rb)
		// println(xa, ya, ra, xb, yb, rb)
		dis := math.Sqrt((xa-xb)*(xa-xb) + (ya-yb)*(ya-yb))
		if dis + ra < rb {
			fmt.Println(-2)
		} else if dis + rb < ra {
			fmt.Println(2)
		} else if dis > rb + ra {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
	}
}
