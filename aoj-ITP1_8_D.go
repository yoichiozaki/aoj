
import (
	"fmt"
	"strings"
)

func main() {
	var s, p string
	fmt.Scan(&s)
	fmt.Scan(&p)
	s = strings.Repeat(s, 2)
	// fmt.Println(p)
	if strings.Index(s, p) == -1 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
