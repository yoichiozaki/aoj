package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

const INF = 0x7FFFFFFF

func main() {
	for scanner.Scan() {
		buf := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(buf[0])
		m, _ := strconv.Atoi(buf[1])
		if n == 0 && m == 0 {
			break
		}
		chair := make([]string, n+1)
		for i := range chair {
			chair[i] = "#"
		}
		chair[n] = "Sentinel"
		for i := 0; i < m; i++ {
			scanner.Scan()
			switch who := scanner.Text(); who {
			case "A": // 左から眺めて行って空いている席があったら即座に座る
				for j := 0; j < n; j++ {
					if chair[j] == "#" {
						chair[j] = "A"
						break
					}
				}
				// log.Println(chair)
				break // 次の人を読み込む
			case "B":
				flag := true
				// 右端からAの隣でない席を探していく
				for j := n - 1; j > -1; j-- {
					if chair[j] == "#" && (j == 0 || chair[j-1] != "A") && chair[j+1] != "A" {
						flag = false
						chair[j] = "B"
						break
					}
				}
				// Aの隣しか空いていないので一番左端の空いている席に座る
				if flag {
					for j := 0; j < n; j++ {
						if chair[j] == "#" {
							chair[j] = "B"
							break
						}
					}
				}
				// log.Println(chair)
				break // 次の人を読み込む
			case "C":
				if i == 0 {
					chair[n/2] = "C" // 先客が誰もいないときは真ん中の椅子に座る
				} else {
					for j := 0; j < n; j++ { // 左から舐めて
						if chair[j] != "#" { // その席が空いていなくて
							if chair[j+1] == "#" { // 右隣が空いているなら座る
								chair[j+1] = "C"
								break
							} else if j != 0 && chair[j-1] == "#" { // 右隣が空いていないので左隣に座る
								chair[j-1] = "C"
								break
							}
						}
					}
				}
				// log.Println(chair)
				break // 次の人を読み込む
			case "D":
				distanceTable := make([]int, n)
				for j := range distanceTable {
					distanceTable[j] = INF
				}
				for j := 0; j < n; j++ {
					if chair[j] != "#" {
						for k := 0; k < n; k++ {
							distanceTable[k] = min(distanceTable[k], abs(j-k))
						}
					}
				}
				// log.Println("distanceTAble:", distanceTable)
				at := 0
				distance := 0
				for j := 0; j < n; j++ {
					if distance < distanceTable[j] {
						at = j
						distance = distanceTable[j]
					}
				}
				chair[at] = "D"
				// log.Println(chair)
				break // 次の人を読み込む
			}
		}
		fmt.Println(strings.Join(chair[:len(chair)-1], ""))
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
