
import "fmt"

func main() {
	var x1, x2, x3, y1, y2, y3, xp, yp float64
	for {
		v, _ := fmt.Scanf("%f %f %f %f %f %f %f %f", &x1, &y1, &x2, &y2, &x3, &y3, &xp, &yp)
		// println(x1, x2, x3, y1, y2, y3, xp, yp)
		if v != 8 {
			break
		}
		ok1 := true
		ok2 := true
		if (x2-x1)*(yp-y1)-(xp-x1)*(y2-y1)<0 {
			ok1 = false
		} else {
			ok2 = false
		}
		if (x3-x2)*(yp-y2)-(xp-x2)*(y3-y2)<0 {
			ok1 = false
		} else {
			ok2 = false
		}
		if (x1-x3)*(yp-y3)-(xp-x3)*(y1-y3)<0 {
			ok1 = false
		} else {
			ok2 = false
		}
		if ok1 || ok2 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
