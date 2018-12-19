// UNCOMPLETED

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
		chairs := make([]string, n)
		for i := range chairs {
			chairs[i] = "#"
		}
		sit := func(who string) {
			if who == "A" {
				chairs[getStringIndex(chairs, "#")] = "A" // A国の人は左から見て行って空いている席を見つけ次第座る
				// log.Println(chairs)
			} else if who == "B" {
				f := true
				for i := n - 1; i > -1; i-- {
					// B国の人はA国人の隣以外で、右端から空いている椅子に座る
					if chairs[i] == "#" && (i == 0 || chairs[i-1] != "A") && (i == n-1 || chairs[i+1] != "A") {
						chairs[i] = "B"
						f = false
						break
					}
				}
				// A国人の隣しか空いていない場合、我慢して左端から空いている椅子に座る
				if f {
					chairs[getStringIndex(chairs, "#")] = "B"
				}
				// log.Println(chairs)
			} else if who == "C" { // C国人は人の隣に座りたい
				for i := 0; i < n; i++ {
					if chairs[i] != "#" {
						if i+1 <= n-1 && chairs[i+1] == "#" {
							chairs[i+1] = "C"
							break
						} else if i-1 >= 0 && chairs[i-1] == "#" {
							chairs[i-1] = "C"
							break
						}
					}
				}
				// log.Println(chairs)
			} else { // who == "D"
				// D国の人は人から離れて座りたい
				// どうしても誰かの隣に座らなければならない場合、その中で一番左側にある椅子に座る
				scores := make([]int, n)
				for i := range scores {
					scores[i] = INF
				}
				for i, v := range chairs {
					if v == "#" {
						scores[i]++
					}
				}
				for i := n - 1; i > -1; i-- {
					tmp := 0
					if chairs[i] == "#" {
						tmp++
					} else {
						tmp = 0
					}
					scores[i] = min(scores[i], tmp)
				}
				chairs[getIntIndex(scores, getMax(scores))] = "D"
				// log.Println(chairs)
			}
		}
		scanner.Scan()
		if first := scanner.Text(); first == "A" || first == "D" {
			chairs[0] = first
		} else if first == "B" {
			chairs[len(chairs)-1] = "B"
		} else {
			chairs[n/2] = "C"
		}
		// log.Println(chairs)
		for i := 0; i < m-1; i++ {
			scanner.Scan()
			sit(scanner.Text())
		}
		fmt.Println(strings.Join(chairs, ""))
	}
}

func getStringIndex(chairs []string, who string) int {
	for i, v := range chairs {
		if v == who {
			return i
		}
	}
	return -1
}

func getIntIndex(scores []int, target int) int {
	for i, v := range scores {
		if v == target {
			return i
		}
	}
	return -1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func getMax(scores []int) int {
	tmp := -1
	for _, v := range scores {
		if tmp < v {
			tmp = v
		}
	}
	return tmp
}
