package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

const MAX_WIDTH = 1500000

func main() {
	// 二分探索で探す
	// 二分探索はソート済みの配列に対して、特定の条件を満たす値を
	// 効率よく選択するアルゴリズム
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), " ")
		m, _ := strconv.Atoi(buf[0])
		n, _ := strconv.Atoi(buf[1])
		if n == 0 && m == 0 {
			break
		}
		books := make([]int, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			books[i], _ = strconv.Atoi(scanner.Text())
		}
		check := func(width int) bool {
			// widthの幅の本棚に順番にbooksを詰めて行ったら
			// 何段になるのかを確認し、それがmを以下であることを確認する
			shelf_index := 1
			position_in_stage := 0
			for i := 0; i < n; i++ {
				// 入れることができない本がある時点でこのwidthは不採用
				if books[i] > width {
					return false
				}
				// 新たに本を追加することができるかをチェックする
				if position_in_stage+books[i] <= width {
					// 新たに本を入れることができたので、そのステージにおける
					// 現在の本の位置を記録
					position_in_stage += books[i]
				} else {
					// 新たな本を入れることができなかったので、一段下のステージに
					// 入れれば良い
					position_in_stage = books[i]
					shelf_index++ // 本棚のステージ数が増えた
				}
			}
			return shelf_index <= m // 条件を満たすかの確認
		}
		low, high := 0, MAX_WIDTH
		for low < high {
			mid := (low + high) / 2
			if check(mid) {
				high = mid
			} else {
				low = mid + 1
			}
			// log.Printf("low: %d, high: %d\n", low, high)
		}
		fmt.Println(low)
	}
}
