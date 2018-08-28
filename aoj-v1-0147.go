
import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

var scanner = bufio.NewScanner(os.Stdin)

type chair struct {
	group_id, rest_time int
}

type info struct {
	group_id, num, wait_sum, before_go int
	fixed                              bool
}

func main() {
	ans := make([]int, 100)
	chair := make([]chair, 17)
	for i := 0; i < 17; i++ {
		chair[i].group_id = -1
	}
	info := make([]info, 100)
	for i := 0; i <= 99; i++ {
		info[i].group_id = i
		info[i].before_go = 5 * i
		info[i].wait_sum = 0
		if i%5 == 1 {
			info[i].num = 5
		} else {
			info[i].num = 2
		}
		info[i].fixed = false
	}
	count := 0
	FLG := 0
	tmp_num := 0
	isOK := false
	waitFLG := false
	for {
		for i := 0; i < 17; i++ {
			if chair[i].group_id != -1 {
				chair[i].rest_time--
				if chair[i].rest_time <= 0 {
					chair[i].group_id = -1
				}
			}
		}
		waitFLG = false
		count = 0
		for i := 0; i < 100; i++ {
			if info[i].fixed {
				continue
			}
			if waitFLG == false && info[i].before_go == 0 {
				FLG = -1
				tmp_num = info[i].num
				for k := 0; k <= 17-tmp_num; k++ {
					isOK = true
					for p := 0; p < tmp_num; p++ {
						if chair[k+p].group_id != -1 {
							isOK = false
							break
						}
					}
					if isOK {
						FLG = k
						break
					}
				}
				if FLG != -1 {
					for k := FLG; k < FLG+tmp_num; k++ {
						chair[k].group_id = info[i].group_id
						chair[k].rest_time = 17*(i%2) + 3*(i%3) + 19
						info[i].fixed = true
					}
				} else {
					waitFLG = true
					info[i].wait_sum++
					count++
				}
			} else {
				if info[i].before_go == 0 {
					info[i].wait_sum++
				} else {
					info[i].before_go--
				}
				count++
			}
		}
		if count == 0 {
			break
		}
	}
	for i := 0; i <= 99; i++ {
		ans[i] = info[i].wait_sum
	}
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		fmt.Println(ans[n])
	}
}

