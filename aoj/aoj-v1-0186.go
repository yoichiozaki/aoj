package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), " ")
		if buf[0] == "0" {
			break
		}
		q1, _ := strconv.Atoi(buf[0]) // 買うべき鶏肉の量の下限(g)
		b, _ := strconv.Atoi(buf[1])  // 予算
		c1, _ := strconv.Atoi(buf[2]) // 会津地鶏の値段(/100g)
		c2, _ := strconv.Atoi(buf[3]) // 普通の鶏肉の値段(/100g)
		q2, _ := strconv.Atoi(buf[4]) // 会津地鶏の買える量の上限(g)
		maxAizuChiken := min(b/c1, q2)
		if maxAizuChiken <= 0 {
			fmt.Println("NA")
			continue
		}
		if c2 >= c1 {
			maxNormalChiken := (b - maxAizuChiken*c1) / c2
			if maxAizuChiken+maxNormalChiken < q1 { // 買い切ったけども必要な鶏肉の下限に達しなかった
				fmt.Println("NA")
			} else {
				fmt.Println(maxAizuChiken, maxNormalChiken)
			}
			continue
		}
		if (b-c1)/c2+1 < q1 { // 会津地鶏を1単位+普通の鶏肉を予算目一杯買ったが買うべき下限を超えなかった
			fmt.Println("NA")
			continue
		}
		// 二分探索
		low, high := 0, maxAizuChiken+1
		for high-low > 1 {
			mid := (low + high) / 2
			if check(mid, b, c1, c2, q1) {
				low = mid
			} else {
				high = mid
			}
		}
		fmt.Println(low, (b-low*c1)/c2)
	}
}

func check(mid, b, c1, c2, q1 int) bool {
	if mid+(b-mid*c1)/c2 < q1 {
		return false
	}
	return true
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
