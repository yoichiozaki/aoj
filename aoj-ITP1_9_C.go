
import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var taroScore, hanakoScore int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var taro, hanako string
		fmt.Scan(&taro, &hanako)
		switch strings.Compare(taro, hanako) {
		case 1: taroScore += 3
		case 0: taroScore, hanakoScore = taroScore+1, hanakoScore+1
		case -1: hanakoScore += 3
		}
	}
	fmt.Printf("%d %d\n", taroScore, hanakoScore)
}
