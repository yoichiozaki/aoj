
import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"math"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	dd := make([][]int, 1000)
	for i := range dd {
		dd[i] = make([]int, 1000)
	}
	num := make([]int,1000)
	inv := make([]int, 1001)
	x := make([]int, 1000)
	y := make([]int, 1000)
	vv := make([]int, 1000)
	nowt := make([]float64, 1000)
	rev := make([]int, 1000)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		for i := 0; i < n; i++ {
			scanner.Scan()
			buf1 := strings.Split(scanner.Text(), " ")
			num[i], _ = strconv.Atoi(buf1[0])
			x[i], _ = strconv.Atoi(buf1[1])
			y[i], _ = strconv.Atoi(buf1[2])
			inv[num[i]] = i
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dd[i][j]=(x[i]-x[j])*(x[i]-x[j])+(y[i]-y[j])*(y[i]-y[j])
			}
		}
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		for i := 0; i < m; i++ {
			scanner.Scan()
			buf2 := strings.Split(scanner.Text(), " ")
			c, _ := strconv.Atoi(buf2[0])
			d, _ := strconv.Atoi(buf2[1])
			c = inv[c]
			d = inv[d]
			for j := 0; j < n; j++ {
				vv[j] = 0
				nowt[j]=999999999
				rev[j]=-1
			}
			nowt[c]=0
			for j := 0; j < n; j++ {
				M := 99999999
				at := 0
				for k := 0; k < n; k++ {
					if vv[k] == 0 && float64(M) > nowt[k] {
						M=int(nowt[k])
						at = k
					}
				}
				if M > 90000000 {
					break
				}
				vv[at]=1
				for k := 0; k < n; k++ {
					if vv[k] == 0 && dd[at][k]<=2500&&nowt[k]>nowt[at]+math.Sqrt(float64(dd[at][k])) {
						nowt[k]=nowt[at]+math.Sqrt(float64(dd[at][k]))
						rev[k] = at
					}
				}
			}
			if vv[d] == 0 {
				fmt.Println("NA")
			} else {
				stack := make([]int, 0)
				V := d
				for V != c {
					stack = append(stack, V)
					V = rev[V]
				}
				fmt.Print(num[c])
				for len(stack) > 0 {
					fmt.Printf(" %d", num[stack[len(stack)-1]])
					stack = stack[:len(stack)-1]
				}
				fmt.Println()
			}
		}
	}
}
