import (
	"fmt"
	"os"
	"strconv"
	"bufio"
)

var sc = bufio.NewScanner(os.Stdin)

func main(){
	for sc.Scan(){
		n := nextInt()
		cnt := 0
		for a:=0; a<=9; a++{
			for b:=0; b<=9; b++{
				for c:=0; c<=9; c++{
					for d:=0; d<=9; d++{
						if a+b+c+d==n{
							cnt++
						}
					}
				}
			}
		}
		fmt.Println(cnt)
	}
}

func next()string{
	//sc.Scan()
	return sc.Text()
}

func nextInt()int{
	a, _ := strconv.Atoi(next())
	return a
}
