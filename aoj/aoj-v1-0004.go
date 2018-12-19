
import "fmt"

func main() {
	var Xs, Ys []float32
	Xs = make([]float32, 0)
	Ys = make([]float32, 0)
	var cnt int
	for {
		var a, b, c, d, e, f, x, y float32
		fmt.Scanf("%f %f %f %f %f %f", &a, &b, &c, &d, &e, &f)
		if a == 0 {
			break
		}
		y = (((a * f) - (d * c)) / ((a * e) - (b * d)))
		x = ((c - (b * y)) / a)
		Xs = append(Xs, x)
		Ys = append(Ys, y)
		cnt++
	}
	for i := 0; i < cnt; i++ {
		fmt.Printf("%.3f %.3f\n", Xs[i], Ys[i])
	}
}
